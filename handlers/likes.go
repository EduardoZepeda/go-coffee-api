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
// @Summary      Return liked shop by user,
// @Description  Return liked shop data by user id
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param id path string true "User id"
// @Param page query int false "Page number"
// @Param size query int false "Size number"
// @Success      200  {array}  models.CoffeeShop
// @Failure      500  {object}  types.ApiError
// @Router       /likes/{id} [get]
func GetLikedCoffeeShops(w http.ResponseWriter, r *http.Request) {
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
	likesByUser := models.LikesByUserRequest{Size: size, Page: page, UserId: params["user_id"]}
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

// Follow a user account godoc
// @Summary      Follow user,
// @Description  Follow a user account using its id
// @Tags         follow
// @Accept       json
// @Produce      json
// @Param request body models.FollowUnfollowRequest true "Follow a user account"
// @Success      201  {object}  models.FollowUnfollowRequest
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /following [post]
func LikeCoffeeShop(w http.ResponseWriter, r *http.Request) {
	var LikeRequest = models.LikeUnlikeCoffeeShopRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&LikeRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must include a UserToId field which is a user Id"}, http.StatusBadRequest)
		return
	}
	userId, err := utils.GetDataFromToken(r, "UserId")
	if err != nil {
		web.Respond(w, types.ApiError{Message: "There was an error with your Authorization header token"}, http.StatusBadRequest)
		return
	}
	// Cast userId as String
	LikeRequest.UserId = userId
	err = repository.LikeCoffeeShop(r.Context(), &LikeRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &LikeRequest, http.StatusCreated)
	return
}

// Unlike a user account godoc
// @Summary      Unfollow user,
// @Description  Unfollow a user account using its id
// @Tags         follow
// @Accept       json
// @Produce      json
// @Param request body models.FollowUnfollowRequest true "Unfollow a user account"
// @Success      204  {object}  models.FollowUnfollowRequest
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /following [delete]
func UnlikeCoffeeShop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := utils.GetDataFromToken(r, "UserId")
	if err != nil {
		web.Respond(w, types.ApiError{Message: "There was an error with your Authorization header token"}, http.StatusBadRequest)
		return
	}
	// Cast userId as String
	var unLikeRequest = models.LikeUnlikeCoffeeShopRequest{UserId: userId, ShopId: params["shop_id"]}
	err = repository.UnlikeCoffeeShop(r.Context(), &unLikeRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusNoContent)
	return
}
