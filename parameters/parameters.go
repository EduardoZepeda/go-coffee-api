package parameters

import (
	"errors"
	"net/http"
	"strconv"
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
