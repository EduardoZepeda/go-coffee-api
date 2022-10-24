package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/golang-jwt/jwt/v4"
)

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
		fmt.Println(err)
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &followRequest, http.StatusAccepted)
	return
}

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
