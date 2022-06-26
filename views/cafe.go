package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EduardoZepeda/go-coffee-api/models"
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
	query := "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop ORDER BY name DESC LIMIT $1 OFFSET $2;"
	db, _ := web.ConnectToDB()
	var cafes []models.Shop
	if err := db.SelectContext(r.Context(), &cafes, query, size, page*size); err != nil {
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
	db, _ := web.ConnectToDB()
	var cafes models.Shop
	query := "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE id = $1;"
	if err := db.Get(&cafes, query, params["id"]); err != nil {
		// If not found {} and 404
		web.Respond(w, struct{}{}, http.StatusNotFound)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}

func CreateCafe(w http.ResponseWriter, r *http.Request) {
	var shopRequest = models.UpsertShop{}
	if err := json.NewDecoder(r.Body).Decode(&shopRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include name, location, address and rating."}, http.StatusBadRequest)
		return
	}
	db, _ := web.ConnectToDB()
	query := "INSERT INTO shops_shop (name, location, address, rating) VALUES ($1, $2, $3, $4);"
	if err := db.MustExec(query, shopRequest.Name, shopRequest.Location, shopRequest.Address, shopRequest.Rating); err != nil {
		web.Respond(w, struct{}{}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, shopRequest, http.StatusCreated)
}

func UpdateCafe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var shopRequest = models.UpsertShop{}
	if err := json.NewDecoder(r.Body).Decode(&shopRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include name, location, address and rating."}, http.StatusBadRequest)
		return
	}
	db, _ := web.ConnectToDB()
	query := "UPDATE shops_shop SET name = $2, location = $3, address = $4, rating = $5 WHERE id = $1;"
	if err := db.MustExec(query, params["id"], shopRequest.Name, shopRequest.Location, shopRequest.Address, shopRequest.Rating); err != nil {
		web.Respond(w, struct{}{}, http.StatusBadRequest)
		return
	}
	web.Respond(w, &shopRequest, http.StatusOK)
}

func DeleteCafe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := web.ConnectToDB()
	query := "DELETE FROM shops_shop WHERE id = $1;"
	if err := db.MustExec(query, params["id"]); err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
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
	db, _ := web.ConnectToDB()
	var cafes []models.Shop
	query := "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE to_tsvector(COALESCE(name, '') || COALESCE(address, '')) @@ plainto_tsquery($1) LIMIT $2 OFFSET $3;"
	if err := db.SelectContext(r.Context(), &cafes, query, params["searchTerm"], size, page*size); err != nil {
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

func GetNearestCafe(w http.ResponseWriter, r *http.Request) {
	var userCoordinates = models.UserCoordinates{}
	if err := json.NewDecoder(r.Body).Decode(&userCoordinates); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must a longitude and a latitude. Example: "}, http.StatusBadRequest)
		return
	}
	db, _ := web.ConnectToDB()
	query := "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop ORDER BY location <-> ST_SetSRID(ST_MakePoint($1, $2), 4326) LIMIT 10;"
	var cafes []models.Shop
	if err := db.SelectContext(r.Context(), &cafes, query, userCoordinates.Latitude, userCoordinates.Longitude); err != nil {
		fmt.Println(err)
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
