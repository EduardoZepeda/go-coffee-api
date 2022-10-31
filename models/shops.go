package models

import (
	"time"

	"github.com/EduardoZepeda/go-coffee-api/types"
)

type CoffeeShop struct {
	ID           string      `db:"id" json:"id,omitempty" swaggerignore:"true"`
	Name         string      `db:"name" json:"name,omitempty"`
	Address      string      `db:"address" json:"address,omitempty"`
	City         string      `db:"city" json:"city,omitempty"`
	Roaster      bool        `db:"roaster" json:"roaster"`
	Location     types.Point `db:"location" json:"location" sql:"type:geometry"`
	Rating       float32     `db:"rating" json:"rating,omitempty"`
	CreatedDate  time.Time   `db:"created_date" json:"created_date,omitempty" swaggerignore:"true"`
	ModifiedDate time.Time   `db:"modified_date" json:"modified_date,omitempty" swaggerignore:"true"`
}
