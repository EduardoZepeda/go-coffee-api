package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/utils"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/golang-jwt/jwt/v4"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest = models.SignUpLoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include an email and a password."}, http.StatusBadRequest)
		return
	}
	user, err := repository.GetUser(r.Context(), loginRequest.Email)
	if err == sql.ErrNoRows {
		web.Respond(w, types.ApiError{Message: "Invalid credentials"}, http.StatusUnauthorized)
		return
	}
	passwordChecker := utils.NewDjangoPassword(user.Password)
	// Verify hashes a password and compares it with the hash passed in when initialized
	passwordsMatched := passwordChecker.VerifyPassword(loginRequest.Password)
	if !passwordsMatched {
		web.Respond(w, types.ApiError{Message: "Invalid credentials"}, http.StatusUnauthorized)
		return
	}
	claims := models.AppClaims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
		},
	}
	//SigningMethodES256 is different than SigningMethodHS256, the later doesn't require a RSA Priv Key as a Signed String
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong in the server"}, http.StatusInternalServerError)
		return
	}
	tokenResponse := models.LoginResponse{
		Token: tokenString,
	}
	web.Respond(w, tokenResponse, http.StatusCreated)
}

func SignupUser(w http.ResponseWriter, r *http.Request) {
	var signupLoginRequest = models.SignUpLoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&signupLoginRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include an email and a password."}, http.StatusBadRequest)
		return
	}
	web.Respond(w, types.ApiError{Message: "Thank you for your interest, new users registration is disabled, because it is handled in a personal way. Please contact admin for a further interview."}, http.StatusUnauthorized)
}
