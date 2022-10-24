package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/utils"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// Login a user godoc
// @Summary      Login a user,
// @Description  Login a user using email and password receive a JWT as a response from a successful login
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body models.LoginRequest true "Login data: email and password"
// @Success      200  {object}  models.LoginResponse
// @Failure      400  {object}  types.ApiError
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest = models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must include an email and a password."}, http.StatusBadRequest)
		return
	}
	user, err := repository.GetUser(r.Context(), loginRequest.Email)
	if err == sql.ErrNoRows {
		web.Respond(w, types.ApiError{Message: "Invalid credentials"}, http.StatusNotFound)
		return
	}
	passwordChecker, err := utils.NewDjangoPassword(user.Password)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Internal server error"}, http.StatusInternalServerError)
		return
	}
	// Verify hashes a password and compares it with the hash passed in when initialized
	passwordsMatched := passwordChecker.VerifyPassword(loginRequest.Password)
	if !passwordsMatched {
		web.Respond(w, types.ApiError{Message: "Invalid credentials"}, http.StatusNotFound)
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

// Register a new user godoc
// @Summary      Register a new user,
// @Description  Register a user using email, username, password and password confirmation
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body models.SignUpRequest true "Login data: email, password and password confirmation"
// @Success      201  {object}  models.EmptyBody
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /user [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var SignUpRequest = models.SignUpRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&SignUpRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid sintax. Request body must include an email, password, passwordConfirmation and username fields."}, http.StatusBadRequest)
		return
	}
	if SignUpRequest.Password != SignUpRequest.PasswordConfirmation {
		web.Respond(w, types.ApiError{Message: "Your password confirmation didn't match your password. Please make sure both are the same."}, http.StatusBadRequest)
		return
	}
	_, err := mail.ParseAddress(SignUpRequest.Email)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Please enter a valid email address"}, http.StatusBadRequest)
		return
	}
	hashedPassword, err := utils.GenerateDjangoHashedPassword(SignUpRequest.Password)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Your password couldn't be processed"}, http.StatusInternalServerError)
		return
	}
	SignUpRequest.HashedPassword = hashedPassword
	err = repository.RegisterUser(r.Context(), &SignUpRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: "Something went wrong with your user registration process."}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, struct{}{}, http.StatusCreated)
	return
}

// Update the current user godoc
// @Summary      Update current user,
// @Description  Update the current user's bio, first name, last name and username
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Param request body models.UpdateUserRequest true "User data: id, bio, firstName, lastName and username"
// @Success      200  {object}  models.UpdateUserRequest
// @Failure      400  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /user/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var UpdateUserRequest = models.UpdateUserRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&UpdateUserRequest); err != nil {
		web.Respond(w, types.ApiError{Message: "Invalid syntax. Request body must include an id, bio, firstName, lastName and username fields."}, http.StatusBadRequest)
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
	// If claims from JWT token and params are differente raise an error
	if params["id"] != claims["userId"] {
		web.Respond(w, types.ApiError{Message: "You can't update other people's data."}, http.StatusBadRequest)
		return
	}
	// Ignore any id coming from user, and assign it to params id
	UpdateUserRequest.Id = params["id"]
	err = repository.UpdateUser(r.Context(), &UpdateUserRequest)
	if err != nil {
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	web.Respond(w, &UpdateUserRequest, http.StatusAccepted)
	return
}

// Get user account data godoc
// @Summary      Get an user account data,
// @Description  Get id, username, email, first name, last name and bio from a user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Success      200  {object}  models.GetUserResponse
// @Failure      404  {object}  types.ApiError
// @Failure      500  {object}  types.ApiError
// @Router       /user/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := repository.GetUserById(r.Context(), params["id"])
	switch err {
	case nil:
		web.Respond(w, user, http.StatusOK)
	case sql.ErrNoRows:
		web.Respond(w, struct{}{}, http.StatusNotFound)
		return
	default:
		web.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
}
