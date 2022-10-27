package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/utils"
	"github.com/gorilla/mux"
)

// Return the list of following users godoc
// @Summary      Return following users,
// @Description  Return following users from a given user Id
// @Tags         follow
// @Accept       json
// @Produce      json
// @Param id path string true "User id"
// @Success      200  {array}  models.GetUserResponse
// @Failure      500  {object}  types.ApiError
// @Router       /following/{id} [get]
func GetUserFollowingAccounts(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		users, err := app.Repo.GetUserFollowing(r.Context(), params["id"])
		if err != nil {
			app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		if len(users) == 0 {
			app.Respond(w, struct{}{}, http.StatusOK)
			return
		}
		app.Respond(w, users, http.StatusOK)
		return
	}
}

// Return the list of user's followers  godoc
// @Summary      Return user's followers,
// @Description  Return user's followers from a given user Id
// @Tags         follow
// @Accept       json
// @Produce      json
// @Param id path string true "User id"
// @Success      200  {array}  models.GetUserResponse
// @Failure      500  {object}  types.ApiError
// @Router       /followers/{id} [get]
func GetUserFollowers(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		users, err := app.Repo.GetUserFollowers(r.Context(), params["id"])
		if err != nil {
			app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		if len(users) == 0 {
			app.Respond(w, struct{}{}, http.StatusOK)
			return
		}
		app.Respond(w, users, http.StatusOK)
		return
	}
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
func FollowUser(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var followRequest = models.FollowUnfollowRequest{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&followRequest); err != nil {
			app.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must include a UserToId field which is a user Id"}, http.StatusBadRequest)
			return
		}
		userId, err := utils.GetDataFromToken(r, "userId")
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an error with your Authorization header token"}, http.StatusBadRequest)
			return
		}
		// Cast userId as String
		followRequest.UserFromId = userId.(string)
		err = app.Repo.FollowUser(r.Context(), &followRequest)
		if err != nil {
			app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, &followRequest, http.StatusCreated)
		return
	}
}

// Unfollow a user account godoc
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
func UnfollowUser(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var unfollowRequest = models.FollowUnfollowRequest{}
		userId, err := utils.GetDataFromToken(r, "userId")
		if err != nil {
			app.Respond(w, types.ApiError{Message: "There was an error with your Authorization header token"}, http.StatusBadRequest)
			return
		}
		// Cast userId as String
		unfollowRequest.UserFromId = userId.(string)
		// Obtain UserToId from url path
		unfollowRequest.UserToId = params["id"]
		err = app.Repo.UnfollowUser(r.Context(), &unfollowRequest)
		if err != nil {
			app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, struct{}{}, http.StatusNoContent)
		return
	}
}
