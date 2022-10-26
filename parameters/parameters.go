package parameters

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/EduardoZepeda/go-coffee-api/models"
)

func GetIntParam(r *http.Request, parameter string, defaultValue uint64) (uint64, error) {
	var err error
	pageStr := r.URL.Query().Get(parameter)
	var page = uint64(defaultValue)
	if pageStr != "" {
		page, err = strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			return page, errors.New("Parameter must be a positive integer. For example: &page=1")
		}
	}
	return page, nil
}

func GetStringParam(r *http.Request, parameter string, defaultValue string) string {
	value := r.URL.Query().Get(parameter)
	if value != "" {
		return value
	}
	return defaultValue
}

func GetLongitudeAndLatitudeTerms(r *http.Request) (*models.UserCoordinates, error) {
	if r.URL.Query().Get("longitude") == "" || r.URL.Query().Get("latitude") == "" {
		return nil, errors.New("Both latitude and longitude must be present as query parameters")
	}
	longitude, err := strconv.ParseFloat(r.URL.Query().Get("longitude"), 32)
	if err != nil {
		return nil, err
	}
	latitude, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 32)
	if err != nil {
		return nil, err
	}
	coordinates := &models.UserCoordinates{
		Longitude: float32(longitude),
		Latitude:  float32(latitude),
	}
	return coordinates, nil
}
