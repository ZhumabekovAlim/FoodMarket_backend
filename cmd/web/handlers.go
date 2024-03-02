package main

import (
	"github.com/gorilla/sessions"
	_ "github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("12345"))

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	products, err := app.product.GetAllProducts()
	if err != nil {
		app.serverError(w, err)
		return
	}

	//id := app.session.Get(r, "authenticatedUserID").(int)
	//user, _ := app.user.Get(id)c
	w.Write(products)
}
