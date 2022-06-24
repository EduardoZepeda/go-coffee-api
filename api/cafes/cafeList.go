package cafes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/EduardoZepeda/go-coffee-api/api/types"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/gorilla/mux"
)

type Shop struct {
	ID           int         `db:"id" json:"id,omitempty"`
	Name         string      `db:"name" json:"name,omitempty"`
	Address      string      `db:"address" json:"address,omitempty"`
	Location     types.Point `sql:"type:geometry"`
	Rating       float32     `db:"rating" json:"rating,omitempty"`
	CreatedDate  time.Time   `db:"created_date" json:"created_date,omitempty"`
	ModifiedDate time.Time   `db:"modified_date" json:"modified_date,omitempty"`
}

func GetCafes(w http.ResponseWriter, r *http.Request) {
	var err error
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Default offset value of 10
	var page = uint64(0)
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
	query := "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop ORDER BY name DESC LIMIT $1 OFFSET $2;"
	db, _ := web.ConnectToDB()
	var cafes []Shop
	if err := db.SelectContext(r.Context(), &cafes, query, size, page*size); err != nil {
		web.Respond(w, cafes, http.StatusInternalServerError)
		return
	}
	if len(cafes) == 0 {
		// if query returns nothing return 404 and []
		web.Respond(w, []int{}, http.StatusNotFound)
		return

	}
	web.Respond(w, cafes, http.StatusOK)
}

func GetCafeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := web.ConnectToDB()
	var cafes Shop
	query := "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE id = $1;"
	if err := db.Get(&cafes, query, params["id"]); err != nil {
		// If not found {} and 404
		web.Respond(w, struct{}{}, http.StatusNotFound)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}

func SearchCafe(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Default offset value of 5
	var page = uint64(0)
	var size = uint64(5)

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
	db, _ := web.ConnectToDB()
	var cafes []Shop
	query := "SELECT id, name, location, address, rating, created_date, modified_date FROM shops_shop WHERE to_tsvector(COALESCE(name, '') || COALESCE(address, '')) @@ plainto_tsquery($1) LIMIT $2 OFFSET $3;"
	if err := db.SelectContext(r.Context(), &cafes, query, params["searchTerm"], size, page*size); err != nil {
		web.Respond(w, &cafes, http.StatusInternalServerError)
		return
	}
	if len(cafes) == 0 {
		// if query returns nothing return 404 and []
		web.Respond(w, []int{}, http.StatusNotFound)
		return
	}
	web.Respond(w, cafes, http.StatusOK)
}
