package main

import "net/http"

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := app.user.GetAllUsers()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(allUsers)
}
