package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/parameters"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"

	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/gorilla/mux"
)

func GetCafes(w http.ResponseWriter, r *http.Request) {
	var err error
	page, err := parameters.GetPage(r)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	size, err := parameters.GetSize(r)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	cafes, err := repository.GetCafes(r.Context(), page, size)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	if len(cafes) == 0 {
		// if query returns nothing return 404 and []
		web.Respond(w, []int{}, http.StatusNotFound)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}

func GetCafeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cafe, err := repository.GetCafeById(r.Context(), params["id"])
	switch err {
	case nil:
		web.Respond(w, cafe, http.StatusOK)
	case sql.ErrNoRows:
		web.Respond(w, struct{}{}, http.StatusNotFound)
		return
	default:
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
}

func CreateCafe(w http.ResponseWriter, r *http.Request) {
	var shopRequest = models.CreateShop{}
	decoder := json.NewDecoder(r.Body)
	// If the JSON in the body has other fields not included in the struct from the parameter, returns an error.
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&shopRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include name, location, address and rating."}, http.StatusBadRequest)
		return
	}
	err := repository.CreateCafe(r.Context(), &shopRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, shopRequest, http.StatusCreated)
}

func UpdateCafe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var shopRequest = models.InsertShop{}
	shopRequest.ID = params["id"]
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&shopRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include name, location, address and rating."}, http.StatusBadRequest)
		return
	}
	err := repository.UpdateCafe(r.Context(), &shopRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &shopRequest, http.StatusOK)
}

func DeleteCafe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repository.DeleteCafe(r.Context(), params["id"])
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusNoContent)
}

func SearchCafe(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)
	page, err := parameters.GetPage(r)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	size, err := parameters.GetSize(r)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	cafes, err := repository.SearchCafe(r.Context(), params["searchTerm"], page, size)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	if len(cafes) == 0 {
		// if query returns nothing return 404 and []
		web.Respond(w, []int{}, http.StatusNotFound)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}

func GetNearestCafes(w http.ResponseWriter, r *http.Request) {
	var userCoordinates = models.UserCoordinates{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&userCoordinates); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must a longitude and a latitude. For example: {'latitude': -103.3668161, 'longitude': 20.6708447}"}, http.StatusBadRequest)
		return
	}
	cafes, err := repository.GetNearestCafes(r.Context(), &userCoordinates)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	if len(cafes) == 0 {
		// if query returns nothing return 404 and []
		web.Respond(w, []int{}, http.StatusNotFound)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}
