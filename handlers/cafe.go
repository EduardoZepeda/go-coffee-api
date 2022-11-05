package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/parameters"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/validator"

	"github.com/gorilla/mux"
)

// GetCoffeeShops godoc
// @Summary      Get a list of coffee shops
// @Description  Get a list of all coffee shop in Guadalajara. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags         coffee shops
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Param search query string false "Search term"
// @Param longitude query float32 false "User longitude"
// @Param latitude query float32 false "User latitude"
// @Success      200  {array}  models.CoffeeShop
// @Failure      404  {object}  []models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops [get]
func GetCoffeeShops(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		page, err := parameters.GetIntParam(r, "page", 0)
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusBadRequest)
			return
		}
		size, err := parameters.GetIntParam(r, "size", 10)
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusBadRequest)
			return
		}
		// If there is search term parameter
		searchTerm := parameters.GetStringParam(r, "search", "")
		if searchTerm != "" {
			cafes, err := app.Repo.SearchCoffeeShops(r.Context(), searchTerm, page, size)
			if err != nil {
				app.Logger.Println(err)
				app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
				return
			}
			if len(cafes) == 0 {
				// if query returns nothing return 404 and []
				app.Respond(w, []int{}, http.StatusNotFound)
				return
			}
			app.Respond(w, cafes, http.StatusOK)
			return
		}
		// if there is latitude and longitude in parameters
		userCoordinates, err := parameters.GetLongitudeAndLatitudeTerms(r)
		if userCoordinates != nil && err == nil {
			cafes, err := app.Repo.GetNearestCoffeeShop(r.Context(), userCoordinates)
			if err != nil {
				app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
				return
			}
			if len(cafes) == 0 {
				// if query returns nothing return 404 and []
				app.Respond(w, []int{}, http.StatusNotFound)
				return
			}
			app.Respond(w, cafes, http.StatusOK)
			return
		}
		// List coffes in a default way
		cafes, err := app.Repo.GetCoffeeShops(r.Context(), page, size)
		if err != nil {
			app.Logger.Println(err)
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		if len(cafes) == 0 {
			// if query returns nothing return 404 and []
			app.Respond(w, []int{}, http.StatusNotFound)
			return
		}
		app.Respond(w, cafes, http.StatusOK)
		return
	}
}

// GetCoffeeShopById godoc
// @Summary      Get a coffee shop by its id
// @Description  Get a specific coffee shop object. Id parameter must be an integer.
// @Tags         coffee shops
// @Accept       json
// @Produce      json
// @Param coffee_shop_id path string true "Coffee Shop ID"
// @Success      200  {object}  models.CoffeeShop
// @Failure      404  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops/{coffee_shop_id} [get]
func GetCoffeeShopById(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		cafe, err := app.Repo.GetCoffeeShopById(r.Context(), params["id"])
		switch err {
		case nil:
			app.Respond(w, cafe, http.StatusOK)
		case sql.ErrNoRows:
			app.Respond(w, struct{}{}, http.StatusNotFound)
			return
		default:
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
	}
}

// CreateCoffeeShop godoc
// @Summary      Create a new coffee shop
// @Description  Create a coffee shop object.
// @Tags         coffee shops
// @Accept       json
// @Produce      json
// @Param request body models.CoffeeShop true "New Coffee Shop data"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Success      201  {object}  models.CoffeeShop
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops [post]
func CreateCoffeeShop(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var coffeeShop = models.CoffeeShop{}
		decoder := json.NewDecoder(r.Body)
		// If the JSON in the body has other fields not included in the struct from the parameter, returns an error.
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&coffeeShop); err != nil {
			app.Respond(w, types.ApiError{Message: "Invalid JSON syntax in body request."}, http.StatusBadRequest)
			return
		}
		v := validator.New()
		if validator.ValidateCoffeeShop(v, &coffeeShop); !v.Valid() {
			app.Respond(w, types.ApiError{Errors: &v.Errors}, http.StatusBadRequest)
			return
		}
		insertedId, err := app.Repo.CreateCoffeeShop(r.Context(), &coffeeShop)
		if err != nil {
			app.Logger.Println(err)
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		coffeeShop.ID = insertedId
		app.Respond(w, coffeeShop, http.StatusCreated)
		return
	}
}

// UpdateCoffeeShop godoc
// @Summary      Update a coffee shop
// @Description  Update a coffee shop object by its Id.
// @Tags         coffee shops
// @Accept       json
// @Produce      json
// @Param request body models.CoffeeShop true "Updated Coffee Shop data"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Param coffee_shop_id path string true "Coffee Shop ID"
// @Success      200  {object}  models.CoffeeShop
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops/{coffee_shop_id} [put]
func UpdateCoffeeShop(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var coffeeShop = models.CoffeeShop{}
		coffeeShop.ID = params["id"]
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&coffeeShop); err != nil {
			app.Respond(w, types.ApiError{Message: "Invalid JSON syntax in body request."}, http.StatusBadRequest)
			return
		}
		v := validator.New()
		if validator.ValidateCoffeeShop(v, &coffeeShop); !v.Valid() {
			app.Respond(w, types.ApiError{Errors: &v.Errors}, http.StatusBadRequest)
			return
		}
		err := app.Repo.UpdateCoffeeShop(r.Context(), &coffeeShop)
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, &coffeeShop, http.StatusOK)
	}
}

// DeleteCoffeeShop godoc
// @Summary      Delete a coffee shop
// @Description  Delete a coffee shop object by its Id.
// @Tags         coffee shops
// @Accept       json
// @Produce      json
// @Param coffee_shop_id path string true "Coffee Shop ID"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Success      204  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops/{coffee_shop_id} [delete]
func DeleteCoffeeShop(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := app.Repo.DeleteCoffeeShop(r.Context(), params["id"])
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, struct{}{}, http.StatusNoContent)
	}
}
