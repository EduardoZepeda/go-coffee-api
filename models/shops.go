package models

import (
	"time"

	"github.com/EduardoZepeda/go-coffee-api/types"
)

type EmptyBody struct{}

type Shop struct {
	ID           string      `db:"id" json:"id,omitempty"`
	Name         string      `db:"name" json:"name,omitempty"`
	Address      string      `db:"address" json:"address,omitempty"`
	Location     types.Point `sql:"type:geometry"`
	Rating       float32     `db:"rating" json:"rating,omitempty"`
	CreatedDate  time.Time   `db:"created_date" json:"created_date,omitempty"`
	ModifiedDate time.Time   `db:"modified_date" json:"modified_date,omitempty"`
}

type CreateShop struct {
	Name     string      `db:"Name" json:"name,omitempty"`
	Address  string      `db:"Adress" json:"address,omitempty"`
	Location types.Point `db:"Location" sql:"type:geometry"`
	Rating   float32     `db:"Rating" json:"rating,omitempty"`
}

type InsertShop struct {
	ID       string      `db:"id" json:"id,omitempty"`
	Name     string      `db:"Name" json:"name,omitempty"`
	Address  string      `db:"Adress" json:"address,omitempty"`
	Location types.Point `db:"Location" sql:"type:geometry"`
	Rating   float32     `db:"Rating" json:"rating,omitempty"`
}
