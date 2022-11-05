package handlers

import (
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/parameters"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/gorilla/mux"
)

// GetCoffeeBagByCoffeeShop godoc
// @Summary      Get a list of coffee bags by coffee shop
// @Description  Get a list of all coffee bags sold by a given coffee shop in Guadalajara. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags         coffee bags by coffee shop
// @Accept       json
// @Produce      json
// @Param id path string true "Coffee Shop ID"
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Success      200  {array}  models.CoffeeBag
// @Failure      404  {object}  []models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops/{id}/coffee-bags [get]
func GetCoffeeBagByCoffeeShop(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		params := mux.Vars(r)
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
		coffeeBagRequest := models.CoffeeBagByShopId{CoffeeShopId: params["id"], Pagination: models.Pagination{Page: page, Size: size}}
		coffeeBags, err := app.Repo.GetCoffeeBagByCoffeeShop(r.Context(), &coffeeBagRequest)
		if err != nil {
			app.Logger.Println(err)
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		if len(coffeeBags) == 0 {
			// if query returns nothing return 404 and []
			app.Respond(w, []int{}, http.StatusNotFound)
			return
		}
		app.Respond(w, coffeeBags, http.StatusOK)
		return
	}
}

// AddCoffeeBagToCoffeeShop godoc
// @Summary      Add a new coffee bag to a coffee shop
// @Description  Add a new coffee bag to a coffee shop by their ids
// @Tags         coffee bags by coffee shop
// @Accept       json
// @Produce      json
// @Param coffee_bag_id path string true "Coffee Bag ID"
// @Param id path string true "Coffee Shop ID"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Success      201  {object}  models.EmptyBody
// @Failure      400  {object}  types.ApiError
// @Failure      404  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops/{id}/coffee-bags/{coffee_bag_id} [post]
func AddCoffeeBagToCoffeeShop(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := app.Repo.AddCoffeeBagToCoffeeShop(r.Context(), params["coffee_bag_id"], params["id"])
		if err != nil {
			app.Logger.Println(err)
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, struct{}{}, http.StatusCreated)
	}
}

// DeleteCoffeeBagFromCoffeeShop godoc
// @Summary      Remove a coffee bag from a coffee shop
// @Description  Remove a coffee bag from a coffee shop using their ids.
// @Tags         coffee bags by coffee shop
// @Accept       json
// @Produce      json
// @Param coffee_bag_id path string true "Coffee Bag ID"
// @Param id path string true "Coffee Shop ID"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Success      204  {object}  models.EmptyBody
// @Failure      500  {object}  types.ApiError
// @Router       /coffee-shops/{id}/coffee-bags/{coffee_bag_id} [delete]
func DeleteCoffeeBagFromCoffeeShop(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := app.Repo.RemoveCoffeeBagFromCoffeeShop(r.Context(), params["coffee_bag_id"], params["id"])
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an internal server error"}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, struct{}{}, http.StatusNoContent)
	}
}
