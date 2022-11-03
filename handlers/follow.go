package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/gorilla/mux"
)

// Return the list of following users godoc
// @Summary      Return following users,
// @Description  Return following users from a given user Id
// @Tags         follows
// @Accept       json
// @Produce      json
// @Param user_id path string true "User id"
// @Param Authorization header string true "With the bearer started."
// @Success      200  {array}  models.GetUserResponse
// @Failure      500  {object}  types.ApiError
// @Router       /following/{user_id} [get]
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
// @Tags         follows
// @Accept       json
// @Produce      json
// @Param user_id path string true "User id"
// @Param Authorization header string true "With the bearer started."
// @Success      200  {array}  models.GetUserResponse
// @Failure      500  {object}  types.ApiError
// @Router       /followers/{user_id} [get]
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
// @Tags         follows
// @Accept       json
// @Produce      json
// @Param request body models.FollowUnfollowRequest true "Follow a user account"
// @Param Authorization header string true "With the bearer started."
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
		ctx := r.Context()
		userId := ctx.Value("userId")
		// Cast userId as String
		followRequest.UserFromId = userId.(string)
		err := app.Repo.FollowUser(ctx, &followRequest)
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
// @Tags         follows
// @Accept       json
// @Produce      json
// @Param request body models.FollowUnfollowRequest true "Unfollow a user account"
// @Param user_id path string true "User id"
// @Param Authorization header string true "With the bearer started."
// @Success      204  {object}  models.EmptyBody
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /following/{user_id} [delete]
func UnfollowUser(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var unfollowRequest = models.FollowUnfollowRequest{}
		ctx := r.Context()
		userId := ctx.Value("userId")
		// Cast userId as String
		unfollowRequest.UserFromId = userId.(string)
		// Obtain UserToId from url path
		unfollowRequest.UserToId = params["id"]
		err := app.Repo.UnfollowUser(ctx, &unfollowRequest)
		if err != nil {
			app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		app.Respond(w, struct{}{}, http.StatusNoContent)
		return
	}
}
