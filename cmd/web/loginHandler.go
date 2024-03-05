package main

//func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
//	body, _ := io.ReadAll(r.Body)
//
//	r.Body = io.NopCloser(bytes.NewBuffer(body))
//
//	id, err := app.user.Authenticate(form.GetUserById("email"), form.GetUserById("password"))
//	if err != nil {
//		if errors.Is(err, models.ErrInvalidCredentials) {
//			form.Errors.Add("generic", "Email or Password is incorrect")
//			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
//		} else {
//			app.serverError(w, err)
//		}
//		return
//	}
//	app.session.Put(r, "authenticatedUserID", id)
//
//	http.Redirect(w, r, "/", http.StatusSeeOther)
//}
