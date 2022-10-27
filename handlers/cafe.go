package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/parameters"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/validator"

	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/gorilla/mux"
)

// GetCoffeeShops godoc
// @Summary      Get a list of coffee shops
// @Description  Get a list of all coffee shop in Guadalajara. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags         coffee shop
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Param searchTerm query string false "Search term"
// @Param longitude query float32 false "User longitude"
// @Param latitude query float32 false "User latitude"
// @Success      200  {array}  models.CoffeeShop
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes [get]
func GetCoffeeShops(w http.ResponseWriter, r *http.Request) {
	var err error
	page, err := parameters.GetIntParam(r, "page", 0)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	size, err := parameters.GetIntParam(r, "size", 10)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	// If there is search term parameter
	searchTerm := parameters.GetStringParam(r, "search", "")
	if searchTerm != "" {
		cafes, err := repository.SearchCoffeeShops(r.Context(), searchTerm, page, size)
		if err != nil {
			log.SetFlags(log.Ldate | log.Lshortfile)
			log.Println(err)
			web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		web.Respond(w, cafes, http.StatusOK)
		return
	}
	// if there is latitude and longitude in parameters
	userCoordinates, err := parameters.GetLongitudeAndLatitudeTerms(r)
	if userCoordinates != nil && err == nil {
		cafes, err := repository.GetNearestCoffeeShop(r.Context(), userCoordinates)
		if err != nil {
			log.SetFlags(log.Ldate | log.Lshortfile)
			log.Println(err)
			web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		web.Respond(w, cafes, http.StatusOK)
		return
	}
	// List coffes in a default way
	cafes, err := repository.GetCoffeeShops(r.Context(), page, size)
	if err != nil {
		log.SetFlags(log.Ldate | log.Lshortfile)
		log.Println(err)
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

// GetCoffeeShopById godoc
// @Summary      Get a coffee shop by its id
// @Description  Get a specific coffee shop object. Id parameter must be an integer.
// @Tags         coffee shop
// @Accept       json
// @Produce      json
// @Param id path string true "Coffee Shop ID"
// @Success      200  {object}  models.CoffeeShop
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes/{id} [get]
func GetCoffeeShopById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cafe, err := repository.GetCoffeeShopById(r.Context(), params["id"])
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

// CreateCoffeeShop godoc
// @Summary      Create a new coffee shop
// @Description  Create a coffee shop object.
// @Tags         coffee shop
// @Accept       json
// @Produce      json
// @Param request body models.CoffeeShop true "New Coffee Shop data"
// @Success      200  {object}  models.CoffeeShop
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes [post]
func CreateCoffeeShop(w http.ResponseWriter, r *http.Request) {
	var coffeeShop = models.CoffeeShop{}
	decoder := json.NewDecoder(r.Body)
	// If the JSON in the body has other fields not included in the struct from the parameter, returns an error.
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&coffeeShop); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid JSON syntax in body request."}, http.StatusBadRequest)
		return
	}
	v := validator.New()
	if validator.ValidateCoffeeShop(v, &coffeeShop); !v.Valid() {
		web.Respond(w, types.ApiError{Errors: &v.Errors}, http.StatusBadRequest)
		return
	}
	err := repository.CreateCoffeeShop(r.Context(), &coffeeShop)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, coffeeShop, http.StatusCreated)
	return
}

// UpdateCoffeeShop godoc
// @Summary      Update a coffee shop
// @Description  Update a coffee shop object by its Id.
// @Tags         coffee shop
// @Accept       json
// @Produce      json
// @Param request body models.CoffeeShop true "Updated Coffee Shop data"
// @Success      200  {object}  models.CoffeeShop
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /cafes/{id} [put]
func UpdateCoffeeShop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var coffeeShop = models.CoffeeShop{}
	coffeeShop.ID = params["id"]
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&coffeeShop); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid JSON syntax in body request."}, http.StatusBadRequest)
		return
	}
	v := validator.New()
	if validator.ValidateCoffeeShop(v, &coffeeShop); !v.Valid() {
		web.Respond(w, types.ApiError{Errors: &v.Errors}, http.StatusBadRequest)
		return
	}
	err := repository.UpdateCoffeeShop(r.Context(), &coffeeShop)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &coffeeShop, http.StatusOK)
}

// DeleteCoffeeShop godoc
// @Summary      Delete a coffee shop
// @Description  Delete a coffee shop object by its Id.
// @Tags         coffee shop
// @Accept       json
// @Produce      json
// @Param id path string true "Coffee Shop ID"
// @Success      204  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /cafes/{id} [delete]
func DeleteCoffeeShop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repository.DeleteCoffeeShop(r.Context(), params["id"])
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusNoContent)
}

// SearchCoffeeShops godoc
// @Summary      Search coffee shops by a given word
// @Description  Search coffee shops by a given word, default number of results are 10
// @Tags         coffee shop, search
// @Accept       json
// @Produce      json
// @Param searchTerm path string true "Search term"
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Success      200 {array}  models.CoffeeShop
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Failure      404  {object}  []models.EmptyBody
// @Router       /cafes/search/{searchTerm} [get]
func SearchCoffeeShops(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)
	page, err := parameters.GetIntParam(r, "page", 0)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	size, err := parameters.GetIntParam(r, "size", 10)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	cafes, err := repository.SearchCoffeeShops(r.Context(), params["searchTerm"], page, size)
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

// NearestCafe godoc
// @Summary      Get a list of the nearest coffee shops
// @Description  Get a list of the user nearest coffee shops in Guadalajara, ordered by distance. It needs user's latitude and longitude as float numbers. Treated as POST to prevent third parties saving users' location into databases.
// @Tags         coffee shop,search
// @Accept       json
// @Produce      json
// @Param request body models.UserCoordinates true "User coordinates (latitude, longitude) in JSON"
// @Success      200 {array}  models.CoffeeShop
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Failure      404  {object}  []models.EmptyBody
// @Router       /cafes/nearest [post]
func GetNearestCoffeeShop(w http.ResponseWriter, r *http.Request) {
	var userCoordinates = models.UserCoordinates{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&userCoordinates); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must a longitude and a latitude. For example: {'latitude': -103.3668161, 'longitude': 20.6708447}"}, http.StatusBadRequest)
		return
	}
	cafes, err := repository.GetNearestCoffeeShop(r.Context(), &userCoordinates)
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
