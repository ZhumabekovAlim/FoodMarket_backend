package main

import (
	"bytes"
	"encoding/json"
	"food_market/pkg/models"
	"io"
	"net/http"
)

func (app *application) createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	createdCategory, err := app.category.CreateCategory(&newCategory)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(createdCategory)
}

func (app *application) updateCategory(w http.ResponseWriter, r *http.Request) {
	var updatedCategory models.Category

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	println(r)
	err := json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.category.UpdateCategory(&updatedCategory)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteCategory(w http.ResponseWriter, r *http.Request) {
	var requestParams struct {
		ID int `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil || requestParams.ID == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.category.DeleteCategory(requestParams.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}