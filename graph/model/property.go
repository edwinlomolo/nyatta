package model

import (
	"time"
)

type NewProperty struct {
	Name       string `json:"name"`
	Town       string `json:"town"`
	PostalCode string `json:"postalCode"`
	CreatedBy  string `json:"createdBy"`
}

type Property struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Town       string     `json:"town"`
	Amenities  []Amenity  `json:"amenities"`
	CreatedBy  string     `json:"createdBy"`
	PostalCode string     `json:"postalCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
