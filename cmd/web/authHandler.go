package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"food_market/pkg/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.user.Insert(newUser.Name, newUser.Email, newUser.Phone, newUser.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	userId, err := app.user.Authenticate(user.Email, user.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.clientError(w, http.StatusBadRequest)
			return
		} else {
			app.serverError(w, err)
			return
		}
	}

	responseUser, err := app.user.GetUserById(strconv.Itoa(userId))
	if app.session.Exists(r, "authenticatedUserID") {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(responseUser)
		if err != nil {
			return
		}
		return
	}
	app.session.Put(r, "authenticatedUserID", userId)
	if err != nil {
		return
	}
	_, err = w.Write(responseUser)
	if err != nil {
		return
	}
}

func (app *application) logOut(w http.ResponseWriter, r *http.Request) {
	if !app.session.Exists(r, "authenticatedUserID") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	app.session.Pop(r, "authenticatedUserID")
	w.WriteHeader(http.StatusOK)
}

func (app *application) testGin(c *gin.Context) {
	user, err := app.user.GetUserById("1")
	if err != nil {
		app.clientError(c.Writer, http.StatusOK)
	}
	c.JSON(http.StatusOK, user)
}
