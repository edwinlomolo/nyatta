package model

import "time"

type NewProperty struct {
	Name       string `json:"name"`
	Town       string `json:"town"`
	PostalCode string `json:"postalCode"`
	Type       string `json:"type"`
	MinPrice   int    `json:"minPrice"`
	MaxPrice   int    `json:"maxPrice"`
	CreatedBy  string `json:"createdBy"`
}

type Property struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Town       string     `json:"town"`
	Type       string     `json:"type"`
	Status     string     `json:"status"`
	MinPrice   int        `json:"minPrice"`
	MaxPrice   int        `json:"maxPrice"`
	Amenities  []Amenity  `json:"amenities"`
	CreatedBy  string     `json:"createdBy"`
	PostalCode string     `json:"postalCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
