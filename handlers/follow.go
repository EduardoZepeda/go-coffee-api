package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// Return the list of following users godoc
// @Summary      Return following users,
// @Description  Return following users from a given user Id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path string true "User id"
// @Success      200  {array}  models.GetUserResponse
// @Failure      500  {object}  types.ApiError
// @Router       /following/{id} [get]
func GetUserFollowingAccounts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	users, err := repository.GetUserFollowing(r.Context(), params["id"])
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		web.Respond(w, struct{}{}, http.StatusOK)
		return
	}
	web.Respond(w, users, http.StatusOK)
	return
}

// Return the list of user's followers  godoc
// @Summary      Return user's followers,
// @Description  Return user's followers from a given user Id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path string true "User id"
// @Success      200  {array}  models.GetUserResponse
// @Failure      500  {object}  types.ApiError
// @Router       /followers/{id} [get]
func GetUserFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	users, err := repository.GetUserFollowers(r.Context(), params["id"])
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		web.Respond(w, struct{}{}, http.StatusOK)
		return
	}
	web.Respond(w, users, http.StatusOK)
	return
}

// Follow a user account godoc
// @Summary      Follow user,
// @Description  Follow a user account using its id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body models.FollowUnfollowRequest true "Follow a user account"
// @Success      201  {object}  models.FollowUnfollowRequest
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /following [post]
func FollowUser(w http.ResponseWriter, r *http.Request) {
	var followRequest = models.FollowUnfollowRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&followRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must include a UserToId field which is a user Id"}, http.StatusBadRequest)
		return
	}
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	// User id is obtained from JWT Token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil && token.Valid {
		web.Respond(w, types.ApiError{Message: "There was an error with your Authorization header token"}, http.StatusBadRequest)
		return
	}
	// Cast userId as String
	followRequest.UserFromId = claims["userId"].(string)
	err = repository.FollowUser(r.Context(), &followRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &followRequest, http.StatusAccepted)
	return
}

// Unfollow a user account godoc
// @Summary      Unfollow user,
// @Description  Unfollow a user account using its id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body models.FollowUnfollowRequest true "Unfollow a user account"
// @Success      204  {object}  models.FollowUnfollowRequest
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /following [delete]
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	var unfollowRequest = models.FollowUnfollowRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&unfollowRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must include a UserToId field which is a user Id"}, http.StatusBadRequest)
		return
	}
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	// User id is obtained from JWT Token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil && token.Valid {
		web.Respond(w, types.ApiError{Message: "There was an error with your Authorization header token"}, http.StatusBadRequest)
		return
	}
	// Cast userId as String
	unfollowRequest.UserFromId = claims["userId"].(string)
	err = repository.UnfollowUser(r.Context(), &unfollowRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusNoContent)
	return
}
