package cafes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/EduardoZepeda/go-coffee-api/web"
)

type Shop struct {
	ID           int       `db:"id" json:"id,omitempty"`
	Name         string    `db:"name" json:"name,omitempty"`
	Address      string    `db:"address" json:"address,omitempty"`
	Rating       float32   `db:"rating" json:"rating,omitempty"`
	CreatedDate  time.Time `db:"created_date" json:"created_date"`
	ModifiedDate time.Time `db:"modified_date" json:"modified_date"`
}

func GetCafes(w http.ResponseWriter, r *http.Request) {
	var err error
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	var page = uint64(0)
	// Default offset value of 10
	var size = uint64(10)

	if pageStr != "" {
		page, err = strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	if sizeStr != "" {
		size, err = strconv.ParseUint(sizeStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	query := fmt.Sprintf("SELECT id, name, address, rating, created_date, modified_date FROM shops_shop ORDER BY name DESC LIMIT %d OFFSET %d;", size, page*size)

	db, _ := web.ConnectToDB()
	var cafes []Shop
	if err := db.SelectContext(r.Context(), &cafes, query); err != nil {
		web.Respond(w, cafes, http.StatusInternalServerError)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}
