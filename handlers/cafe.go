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

// GetCafes godoc
// @Summary      Get a list of coffee shops
// @Description  Get a list of all coffee shop in Guadalajara. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags         cafes
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Success      200  {array}  models.Shop
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes [get]
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

// GetCafeById godoc
// @Summary      Get a new coffee shop by its id
// @Description  Get a specific coffee shop object. Id parameter must be an integer.
// @Tags         cafe
// @Accept       json
// @Produce      json
// @Param id path string true "Coffee Shop ID"
// @Success      200  {object}  models.Shop
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes/{id} [get]
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

// CreateCafe godoc
// @Summary      Create a new coffee shop
// @Description  Create a coffee shop object.
// @Tags         cafe
// @Accept       json
// @Produce      json
// @Param request body models.CreateShop true "New Coffee Shop data"
// @Success      200  {object}  models.CreateShop
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes [post]
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

// UpdateCafe godoc
// @Summary      Update a coffee shop
// @Description  Update a coffee shop object.
// @Tags         cafe
// @Accept       json
// @Produce      json
// @Param request body models.InsertShop true "Updated Coffee Shop data"
// @Success      200  {object}  models.InsertShop
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes/{id} [put]
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

// DeleteCafe godoc
// @Summary      Delete a coffee shop
// @Description  Delete a coffee shop object.
// @Tags         cafe
// @Accept       json
// @Produce      json
// @Param id path string true "Coffee Shop ID"
// @Success      204  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /cafes/{id} [delete]
func DeleteCafe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repository.DeleteCafe(r.Context(), params["id"])
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusNoContent)
}

// SearchCafe godoc
// @Summary      Search a coffee shop by a given word
// @Description  Search a coffee shop by a given word
// @Tags         cafe,search
// @Accept       json
// @Produce      json
// @Param searchTerm path string true "Search term"
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Success      200 {array}  models.Shop
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Failure      404  {object}  []models.EmptyBody
// @Router       /cafes/search/{searchTerm} [get]
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

// NearestCafe godoc
// @Summary      Get a list of the nearest coffee shops
// @Description  Get a list of the user nearest coffee shops in Guadalajara, ordered by distance. It needs user's latitude and longitude as float numbers. Treated as POST to prevent third parties to save users' location into databases.
// @Tags         cafe,search
// @Accept       json
// @Produce      json
// @Param request body models.UserCoordinates true "User coordinates (latitude, longitude) in JSON"
// @Success      200 {array}  models.Shop
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Failure      404  {object}  []models.EmptyBody
// @Router       /cafes/nearest [post]
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
