package model

type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
