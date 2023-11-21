package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID   `json:"id"`
	FirstName        string      `json:"first_name"`
	LastName         string      `json:"last_name"`
	Phone            string      `json:"phone"`
	IsLandlord       bool        `json:"is_landlord"`
	Onboarding       bool        `json:"onboarding"`
	Email            string      `json:"email"`
	SubscribeRetries int32       `json:"subscribe_retries"`
	Avatar           *AnyUpload  `json:"avatar"`
	Properties       []*Property `json:"properties"`
	CreatedAt        *time.Time  `json:"created_at"`
	UpdatedAt        *time.Time  `json:"updated_at"`
}

type SignIn struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
