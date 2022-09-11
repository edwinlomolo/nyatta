package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	FirstName string     `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string     `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string     `gorm:"uniqueIndex" json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UUID default id for user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewRandom()
	tx.Statement.SetColumn("id", uuid.String())
	return
}
