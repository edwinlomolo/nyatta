package model

import "time"

type Amenity struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Provider  string     `json:"provider"`
	Category  string     `json:"category"`
	UnitID    string     `json:"property_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
