package model

import (
	"time"

	"github.com/rs/xid"
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

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Code int    `json:"code"`
	Err  string `json:"error,omitempty"`
}

// Assign default id for user during create
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := xid.New()
	tx.Statement.SetColumn("id", id.String())
	return
}
