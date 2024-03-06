package main

import (
	"bytes"
	"encoding/json"
	"food_market/pkg/models"
	"io"
	"net/http"
	"strconv"
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

	w.WriteHeader(http.StatusCreated)
	w.Write(createdProduct)
}

func (app *application) products(w http.ResponseWriter, r *http.Request) {
	products, err := app.product.GetAllProducts()
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Write(products)
}

func (app *application) productsById(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get(":id")

	productResponse, err := app.product.GetProductById(productId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(productResponse)
}

func (app *application) updateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct models.Product
	productId, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || productId < 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.product.UpdateProduct(&updatedProduct, productId)
	if err != nil {
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || productId < 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.product.DeleteProduct(productId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
