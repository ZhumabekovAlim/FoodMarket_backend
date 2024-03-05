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
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) (*application, error) {
	var user models.User

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		//return sessions.Session{}, user, err
	}

	//id, err := app.user.Authenticate(form.GetUserById("email"), form.GetUserById("password"))
	id, err := app.user.Authenticate(user.Email, user.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.clientError(w, http.StatusBadRequest)
			//return sessions.Session{}, user, err
			return &application{
				session: nil,
			}, err
		} else {
			app.serverError(w, err)
		}
		//return sessions.Session{}, user, err
	}
	app.session.Put(r, "authenticatedUserID", id)
	// Redirect the user to the create snippet page.
	//http.Redirect(w, r, "/", http.StatusSeeOther)

	//return app.session, user, nil
	return &application{
			session: app.session,
			user:    app.user,
		},
		nil
}

func loginUserHandler(app *application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Внутри этой функции вы вызываете метод loginUser вашей структуры app.
		result, err := app.loginUser(w, r)
		// Обработка результата и ошибок...
		if err != nil {
			return
		}
		_ = result
	}
}
