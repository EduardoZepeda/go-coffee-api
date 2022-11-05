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

// GetCoffeeBags godoc
// @Summary      Get a list of coffee bags
// @Description  Get a list of all coffee bags in Guadalajara. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags         coffee bags
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Success      200  {array}  models.CoffeeBag
// @Failure      404  {object}  []models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-bags [get]
func GetCoffeeBags(app *application.App) http.HandlerFunc {
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
		// List coffes bags in a default way
		coffeeBagRequests := models.CoffeeBagsList{Pagination: models.Pagination{Page: page, Size: size}}
		cafes, err := app.Repo.GetCoffeeBags(r.Context(), coffeeBagRequests)
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

// GetCoffeeBagById godoc
// @Summary      Get a coffee bag by its id
// @Description  Get a specific coffee bag object. Id parameter must be an integer.
// @Tags         coffee bags
// @Accept       json
// @Produce      json
// @Param coffee_bag_id path string true "Coffee Bag ID"
// @Success      200  {object}  models.CoffeeBag
// @Failure      404  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-bags/{coffee_bag_id} [get]
func GetCoffeeBagById(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		coffeeBag, err := app.Repo.GetCoffeeBagById(r.Context(), params["id"])
		switch err {
		case nil:
			app.Respond(w, coffeeBag, http.StatusOK)
			return
		case sql.ErrNoRows:
			app.Respond(w, struct{}{}, http.StatusNotFound)
			return
		default:
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
	}
}

// CreateCoffeeBag godoc
// @Summary      Create a new coffee bag
// @Description  Create a coffee bag object.
// @Tags         coffee bags
// @Accept       json
// @Produce      json
// @Param request body models.CoffeeBag true "New Coffee Bag data"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Success      201  {object}  models.CoffeeBag
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-bags [post]
func CreateCoffeeBag(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var coffeeBag = models.CoffeeBag{}
		decoder := json.NewDecoder(r.Body)
		// If the JSON in the body has other fields not included in the struct from the parameter, returns an error.
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&coffeeBag); err != nil {
			app.Respond(w, types.ApiError{Message: "Invalid JSON syntax in body request."}, http.StatusBadRequest)
			return
		}
		v := validator.New()
		if validator.ValidateCoffeeBag(v, &coffeeBag); !v.Valid() {
			app.Respond(w, types.ApiError{Errors: &v.Errors}, http.StatusBadRequest)
			return
		}
		insertedCoffeeBag, err := app.Repo.CreateCoffeeBag(r.Context(), &coffeeBag)
		if err != nil {
			app.Logger.Println(err)
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, insertedCoffeeBag, http.StatusCreated)
		return
	}
}

// UpdateCoffeeBag godoc
// @Summary      Update a coffee bag
// @Description  Update a coffee bag object by its Id.
// @Tags         coffee bags
// @Accept       json
// @Produce      json
// @Param request body models.CoffeeBag true "Updated Coffee Bag data"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Param coffee_bag_id path string true "Coffee Bag ID"
// @Success      200  {object}  models.CoffeeBag
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-bags/{coffee_bag_id} [put]
func UpdateCoffeeBag(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var coffeeBag = models.CoffeeBag{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&coffeeBag); err != nil {
			app.Respond(w, types.ApiError{Message: "Invalid JSON syntax in body request."}, http.StatusBadRequest)
			return
		}
		coffeeBag.ID = params["id"]
		v := validator.New()
		if validator.ValidateCoffeeBag(v, &coffeeBag); !v.Valid() {
			app.Respond(w, types.ApiError{Errors: &v.Errors}, http.StatusBadRequest)
			return
		}
		_, err := app.Repo.UpdateCoffeeBag(r.Context(), &coffeeBag)
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, &coffeeBag, http.StatusOK)
	}
}

// DeleteCoffeeBag godoc
// @Summary      Delete a coffee bag
// @Description  Delete a coffee bag object by its Id.
// @Tags         coffee bags
// @Accept       json
// @Produce      json
// @Param coffee_bag_id path string true "Coffee Bag ID"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Success      204  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-bags/{coffee_bag_id} [delete]
func DeleteCoffeeBag(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := app.Repo.DeleteCoffeeBag(r.Context(), params["id"])
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, struct{}{}, http.StatusNoContent)
	}
}
