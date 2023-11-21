package model

type SignIn struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
