package models

import (
	"time"

	"github.com/EduardoZepeda/go-coffee-api/types"
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

type CreateShop struct {
	Name     string      `db:"name" json:"name,omitempty"`
	Address  string      `db:"address" json:"address,omitempty"`
	Location types.Point `sql:"type:geometry"`
	Rating   float32     `db:"rating" json:"rating,omitempty"`
}

type InsertShop struct {
	ID       string      `db:"id" json:"id,omitempty"`
	Name     string      `db:"name" json:"name,omitempty"`
	Address  string      `db:"address" json:"address,omitempty"`
	Location types.Point `sql:"type:geometry"`
	Rating   float32     `db:"rating" json:"rating,omitempty"`
}
