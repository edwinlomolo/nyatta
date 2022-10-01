package model

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

type AmenityInput struct {
	Name       string `json:"name"`
	Provider   string `json:"provider"`
	PropertyID string `json:"property_id"`
}

type Amenity struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"type:varchar(100);not null;index" json:"name"`
	Provider   string     `gorm:"type:varchar(100);index" json:"provider"`
	PropertyID string     `json:"property_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

// Assign default random id for property
func (a *Amenity) BeforeCreate(tx *gorm.DB) (err error) {
	newId := fmt.Sprintf("N%v", xid.New().String())
	tx.Statement.SetColumn("id", newId)
	return
}
