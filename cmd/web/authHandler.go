package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"food_market/pkg/models"
	"io"
	"net/http"
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

	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
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
	app.session.Put(r, "authenticatedUserID", userId)
	convertedUser, err := json.Marshal(user)
	if err != nil {
		return
	}
	w.Write(convertedUser)
}
