package main

import (
	"bytes"
	"encoding/json"
	"food_market/pkg/models"
	"io"
	"net/http"
	"strconv"
)

func (app *application) profile(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get(":id")
	if !app.session.Exists(r, "authenticatedUserID") {
		app.clientError(w, http.StatusUnauthorized)
		return
	}
	authorizedId := app.session.GetInt(r, "authenticatedUserID")
	userIdToCompare, err := strconv.Atoi(userId)
	if authorizedId != userIdToCompare {
		app.clientError(w, http.StatusForbidden)
		return
	}

	user, err := app.user.GetUserById(userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(user)
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	userId := r.URL.Query().Get(":id")

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		app.serverError(w, err)
		return
	}

	user, err := app.user.GetUserById(userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(user)
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get(":id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	app.user.DeleteUserById(id)
	w.WriteHeader(http.StatusNoContent)
}
