package parameters

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/EduardoZepeda/go-coffee-api/models"
)

func GetPage(r *http.Request) (uint64, error) {
	pageStr := r.URL.Query().Get("page")
	var page = uint64(0)
	if pageStr != "" {
		page, err := strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			return page, errors.New("Page parameter must be a positive integer. For example: &page=1")
		}
	}
	return page, nil
}

func GetSize(r *http.Request) (uint64, error) {
	sizeStr := r.URL.Query().Get("size")
	// Default offset value of 10
	var size = uint64(10)
	if sizeStr != "" {
		size, err := strconv.ParseUint(sizeStr, 10, 64)
		if err != nil {
			return size, errors.New("Size parameter must be a positive integer. For example: &size=5")
		}
	}
	return size, nil
}

func GetSearchTerm(r *http.Request) string {
	return r.URL.Query().Get("search")
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
