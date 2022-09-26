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
}

type Property struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"type:varchar(100);not null" json:"name"`
	Town       string     `gorm:"type:varchar(100):not null;uniqueIndex" json:"town"`
	PostalCode string     `gorm:"type:varchar(6);not null;uniqueIndex" json:"postalCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

// Assign default id for property during create
func (p *Property) BeforeCreate(tx *gorm.DB) (err error) {
	newId := fmt.Sprintf("N%v", xid.New().String())
	tx.Statement.SetColumn("id", newId)
	return
}
