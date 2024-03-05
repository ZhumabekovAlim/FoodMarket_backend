package main

import (
	"net/http"
)

func (app *application) profile(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get(":id")

	user, err := app.user.GetUserById(userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(user)
}
