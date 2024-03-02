package main

import (
	"bytes"
	"encoding/json"
	"food_market/pkg/models"
	"io"
	"net/http"
)

func (app *application) createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	createdProduct, err := app.product.CreateProduct(&newProduct)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(createdProduct)
}

func (app *application) products(w http.ResponseWriter, r *http.Request) {
	products, err := app.product.GetAllProducts()
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(products)
}

func (app *application) updateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct models.Product

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	println(r)
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.product.UpdateProduct(&updatedProduct)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	var requestParams struct {
		ID int `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil || requestParams.ID == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.product.DeleteProduct(requestParams.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
