package model

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

type NewProperty struct {
	Name       string `json:"name"`
	Town       string `json:"town"`
	PostalCode string `json:"postalCode"`
	CreatedBy  string `json:"createdBy"`
}

type Property struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"type:varchar(100);not null" json:"name"`
	Town       string     `gorm:"type:varchar(100);not null;index" json:"town"`
	Amenities  []Amenity  `json:"amenities"`
	CreatedBy  string     `json:"createdBy"`
	PostalCode string     `gorm:"type:varchar(6);not null;index" json:"postalCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

// Assign default random id for property during create
func (p *Property) BeforeCreate(tx *gorm.DB) (err error) {
	newId := fmt.Sprintf("P%v", xid.New().String())
	tx.Statement.SetColumn("id", newId)
	return
}
