package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/parameters"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/utils"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/gorilla/mux"
)

// Return the list of likes by user id godoc
// @Summary      Return liked coffee shops by user,
// @Description  Return liked coffee shops data by user id
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param user query int false "User id"
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Success      200  {array}  models.CoffeeShop
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /likes [get]
func GetLikedCoffeeShops(w http.ResponseWriter, r *http.Request) {
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
	currentUserId, err := utils.GetDataFromToken(r, "userId")
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	userId := parameters.GetStringParam(r, "user", currentUserId.(string))
	likesByUser := models.LikesByUserRequest{Size: size, Page: page, UserId: userId}
	// If there is search term parameter
	shops, err := repository.GetLikedCoffeeShops(r.Context(), &likesByUser)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	if len(shops) == 0 {
		web.Respond(w, struct{}{}, http.StatusOK)
		return
	}
	web.Respond(w, shops, http.StatusOK)
	return
}

// Like a coffee shop godoc
// @Summary      Like a coffee shop
// @Description  Like a coffee shop
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param request body models.LikeUnlikeCoffeeShopRequest true "Like a coffee shop"
// @Success      201  {object}  models.LikeUnlikeCoffeeShopRequest
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /likes [post]
func LikeCoffeeShop(w http.ResponseWriter, r *http.Request) {
	var LikeRequest = models.LikeUnlikeCoffeeShopRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&LikeRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must include a UserToId field which is a user Id"}, http.StatusBadRequest)
		return
	}
	userId, err := utils.GetDataFromToken(r, "userId")
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	// Cast userId as String
	LikeRequest.UserId = userId.(string)
	err = repository.LikeCoffeeShop(r.Context(), &LikeRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &LikeRequest, http.StatusCreated)
	return
}

// Unlike a coffee shop godoc
// @Summary      Unlike a coffee shop
// @Description  Unlike a coffee shop
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param request body models.LikeUnlikeCoffeeShopRequest true "Unlike a coffee shop"
// @Success      201  {object}  models.LikeUnlikeCoffeeShopRequest
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /likes [delete]
func UnlikeCoffeeShop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := utils.GetDataFromToken(r, "userId")
	if err != nil {
		web.Respond(w, types.ApiError{Message: "There was an error with your Authorization header token"}, http.StatusBadRequest)
		return
	}
	// Cast userId as String
	var unLikeRequest = models.LikeUnlikeCoffeeShopRequest{UserId: userId.(string), ShopId: params["shop_id"]}
	err = repository.UnlikeCoffeeShop(r.Context(), &unLikeRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusNoContent)
	return
}
