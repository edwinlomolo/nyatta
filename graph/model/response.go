package model

type Response struct {
	Code int    `json:"code"`
	Err  string `json:"error,omitempty"`
}
