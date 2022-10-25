package utils

import (
	"errors"
	"net/http"
	"strings"
)

func GetTokenFromAuthHeader(r *http.Request) (string, error) {
	AuthHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	token := strings.Split(AuthHeader, " ")
	if (len(token)) != 2 {
		return "", errors.New("Token is not in the format: <Bearer Token>")
	}
	return token[1], nil
}
