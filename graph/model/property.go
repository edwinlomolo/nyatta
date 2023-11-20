package model

import (
	"time"

	"github.com/google/uuid"
)

type NewProperty struct {
	Name      string    `json:"name"`
	Location  *GpsInput `json:"location"`
	Type      string    `json:"type"`
	Thumbnail string    `json:"thumbnail"`
}

type Property struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Location    *Gps            `json:"location"`
	Units       []*PropertyUnit `json:"property_units"`
	Thumbnail   *AnyUpload      `json:"thumbnail"`
	Caretaker   *Caretaker      `json:"caretaker"`
	CaretakerID *uuid.UUID      `json:"caretakerId"`
	Owner       *User           `json:"owner"`
	CreatedBy   uuid.UUID       `json:"createdBy"`
	CreatedAt   *time.Time      `json:"createdAt"`
	UpdatedAt   *time.Time      `json:"updatedAt"`
}
