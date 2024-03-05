package main

import (
	_ "github.com/gorilla/sessions"
	"net/http"
)

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

	w.Write(products)
}
