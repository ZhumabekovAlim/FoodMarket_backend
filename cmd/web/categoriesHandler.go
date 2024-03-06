package main

import (
	"bytes"
	"encoding/json"
	"food_market/pkg/models"
	"io"
	"net/http"
	"strconv"
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

func (app *application) categories(w http.ResponseWriter, r *http.Request) {
	categories, err := app.category.GetAllCategories()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(categories)
}

func (app *application) categoryById(w http.ResponseWriter, r *http.Request) {
	categoryId := r.URL.Query().Get(":id")

	categoryResponse, err := app.category.GetCategoryById(categoryId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(categoryResponse)
}

func (app *application) updateCategory(w http.ResponseWriter, r *http.Request) {
	var updatedCategory models.Category
	categoryId, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || categoryId < 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	println(r)
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.category.UpdateCategory(&updatedCategory, categoryId)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteCategory(w http.ResponseWriter, r *http.Request) {

	categoryId, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || categoryId == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.category.DeleteCategory(categoryId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
