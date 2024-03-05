package main

import (
	"net/http"
)

func (app *application) profile(w http.ResponseWriter, r *http.Request) {

	user, err := app.user.GetUserById(1)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	app.render(w, r, "profile.page.tmpl", &templateData{User: user})
}
