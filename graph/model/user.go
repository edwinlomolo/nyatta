package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"uniqueIndex"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewRandom()
	tx.Statement.SetColumn("id", uuid.String())
	return
}
