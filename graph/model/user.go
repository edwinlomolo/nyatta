package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string     `json:"id";gorm:"primaryKey"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `gorm:"uniqueIndex"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewRandom()
	tx.Statement.SetColumn("id", uuid.String())
	return
}
