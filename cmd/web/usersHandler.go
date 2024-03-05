package main

import (
	"bytes"
	"encoding/json"
	"food_market/pkg/models"
	"io"
	"net/http"
)

func (app *application) profile(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get(":id")

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
