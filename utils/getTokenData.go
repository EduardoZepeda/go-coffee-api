package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func contains(list []string, item string) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}

func GetTokenFromAuthHeader(r *http.Request) (string, error) {
	AuthHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	// Empty Auth header
	if AuthHeader == "" {
		return "", errors.New("Authorization header not provided")
	}
	token := strings.Split(AuthHeader, " ")
	if (len(token)) != 2 {
		return "", errors.New("Authorization header Token is not present in a valid format: <Bearer Token>")
	}
	return token[1], nil
}

func GetDataFromToken(r *http.Request, data string) (interface{}, error) {
	VALID_TOKEN_KEYS := []string{"userId", "isStaff"}
	if !contains(VALID_TOKEN_KEYS, data) {
		return nil, errors.New(fmt.Sprintf("JWT Token doesn't contain the %s claim", data))
	}
	tokenString, err := GetTokenFromAuthHeader(r)
	if err != nil {
		return nil, err
	}
	// User id is obtained from JWT Token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil && token.Valid {
		return nil, err
	}
	return claims[data], nil
}
