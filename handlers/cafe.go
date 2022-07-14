package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"

	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/gorilla/mux"
)

func GetCafes(w http.ResponseWriter, r *http.Request) {
	var err error
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Default offset value of 10
	var page = uint64(0)
	var size = uint64(10)

	if pageStr != "" {
		page, err = strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			web.Respond(w, types.ApiError{Message: "Page parameter must be a positive integer. For example: &page=1"}, http.StatusBadRequest)
			return
		}
	}
	if sizeStr != "" {
		size, err = strconv.ParseUint(sizeStr, 10, 64)
		if err != nil {
			web.Respond(w, types.ApiError{Message: "Size parameter must be a positive integer. For example: &size=5"}, http.StatusBadRequest)
			return
		}
	}
	cafes, err := repository.GetCafes(r.Context(), page, size)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
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
	if err := json.NewDecoder(r.Body).Decode(&shopRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include name, location, address and rating."}, http.StatusBadRequest)
		return
	}
	err := repository.CreateCafe(r.Context(), &shopRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, shopRequest, http.StatusCreated)
}

func UpdateCafe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var shopRequest = models.InsertShop{}
	shopRequest.ID = params["id"]
	if err := json.NewDecoder(r.Body).Decode(&shopRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include name, location, address and rating."}, http.StatusBadRequest)
		return
	}
	err := repository.UpdateCafe(r.Context(), &shopRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &shopRequest, http.StatusOK)
}

func DeleteCafe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repository.DeleteCafe(r.Context(), params["id"])
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusNoContent)
}

func SearchCafe(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Default offset value of 5
	var page = uint64(0)
	var size = uint64(5)

	if pageStr != "" {
		page, err = strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			web.Respond(w, types.ApiError{Message: "Page parameter must be a positive integer. For example: &page=1"}, http.StatusBadRequest)
			return
		}
	}
	if sizeStr != "" {
		size, err = strconv.ParseUint(sizeStr, 10, 64)
		if err != nil {
			web.Respond(w, types.ApiError{Message: "Size parameter must be a positive integer. For example: &size=5"}, http.StatusBadRequest)
			return
		}
	}
	cafes, err := repository.SearchCafe(r.Context(), params["searchTerm"], page, size)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
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
	if err := json.NewDecoder(r.Body).Decode(&userCoordinates); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must a longitude and a latitude. For example: {'latitude': -103.3668161, 'longitude': 20.6708447}"}, http.StatusBadRequest)
		return
	}
	cafes, err := repository.GetNearestCafes(r.Context(), &userCoordinates)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	if len(cafes) == 0 {
		// if query returns nothing return 404 and []
		web.Respond(w, []int{}, http.StatusNotFound)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}
