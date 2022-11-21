package model

import "time"

type AmenityInput struct {
	Name       string `json:"name"`
	Provider   string `json:"provider"`
	PropertyID string `json:"property_id"`
}

type Amenity struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Provider   string     `json:"provider"`
	PropertyID string     `json:"property_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
