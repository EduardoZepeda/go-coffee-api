package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func GetTokenFromAuthHeader(r *http.Request) (string, error) {
	AuthHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	token := strings.Split(AuthHeader, " ")
	if (len(token)) != 2 {
		return "", errors.New("Token is not in the format: <Bearer Token>")
	}
	return token[1], nil
}

func GetDataFromToken(r *http.Request, data string) (string, error) {
	tokenString, err := GetTokenFromAuthHeader(r)
	if err != nil {
		return "", err
	}
	// User id is obtained from JWT Token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil && token.Valid {
		return "", err
	}
	return claims[data].(string), nil
}
