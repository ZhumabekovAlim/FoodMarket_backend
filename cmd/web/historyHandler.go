package main

//
//import (
//	"bytes"
//	"encoding/json"
//	"food_market/pkg/models"
//	"io"
//	"net/http"
//)
//
//func (app *application) createHistory(w http.ResponseWriter, r *http.Request) {
//	var newHistory models.OrderHistory
//
//	body, _ := io.ReadAll(r.Body)
//	r.Body = io.NopCloser(bytes.NewBuffer(body))
//
//	err := json.NewDecoder(r.Body).Decode(&newHistory)
//	if err != nil {
//		app.clientError(w, http.StatusBadRequest)
//		return
//	}
//
//	createdHistory, err := app.history.CreateHistory(&newHistory)
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	w.Write(createdHistory)
//}
//
//func (app *application) historiesByUserId(w http.ResponseWriter, r *http.Request) {
//	var updatedHistory models.OrderHistory
//
//	body, _ := io.ReadAll(r.Body)
//	r.Body = io.NopCloser(bytes.NewBuffer(body))
//	println(r)
//	err := json.NewDecoder(r.Body).Decode(&updatedHistory)
//	if err != nil {
//		app.clientError(w, http.StatusBadRequest)
//		return
//	}
//	histories, err := app.history.GetHistoryByUserId(&updatedHistory)
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//
//	w.Write(histories)
//}
//
//func (app *application) histories(w http.ResponseWriter, r *http.Request) {
//	histories, err := app.history.GetAllHistory()
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//
//	w.Write(histories)
//}
//
//func (app *application) historyById(w http.ResponseWriter, r *http.Request) {
//	historyId := r.URL.Query().Get(":id")
//
//	history, err := app.history.GetHistoryById(historyId)
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//
//	w.Write(history)
//}
//
//func (app *application) updateHistory(w http.ResponseWriter, r *http.Request) {
//	var updatedHistory models.OrderHistory
//
//	body, _ := io.ReadAll(r.Body)
//	r.Body = io.NopCloser(bytes.NewBuffer(body))
//	println(r)
//	err := json.NewDecoder(r.Body).Decode(&updatedHistory)
//	if err != nil {
//		app.clientError(w, http.StatusBadRequest)
//		return
//	}
//
//	err = app.history.UpdateHistory(&updatedHistory)
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
//
//func (app *application) deleteHistory(w http.ResponseWriter, r *http.Request) {
//	var requestParams struct {
//		ID int `json:"id"`
//	}
//
//	err := json.NewDecoder(r.Body).Decode(&requestParams)
//	if err != nil || requestParams.ID == 0 {
//		app.clientError(w, http.StatusBadRequest)
//		return
//	}
//
//	err = app.history.DeleteHistory(requestParams.ID)
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
