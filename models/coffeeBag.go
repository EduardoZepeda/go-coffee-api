package models

type CoffeeBag struct {
	ID      string `json:"id,omitempty" swaggerignore:"true"`
	Brand   string `json:"brand,omitempty"`
	Origin  string `json:"origin,omitempty"`
	Species string `json:"species,omitempty"`
}

type CoffeeBagsList struct {
	Pagination
}
