package model

import "time"

type User struct {
	ID         string       `json:"id"`
	FirstName  string       `json:"first_name"`
	LastName   string       `json:"last_name"`
	Phone      string       `json:"phone"`
	IsLandlord bool         `json:"is_landlord"`
	Onboarding bool         `json:"onboarding"`
	Email      string       `json:"email"`
	Uploads    []*AnyUpload `json:"uploads"`
	Properties []*Property  `json:"properties"`
	CreatedAt  *time.Time   `json:"created_at"`
	UpdatedAt  *time.Time   `json:"updated_at"`
}

type SignIn struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
