package config

import "time"

// JWT - jwt authentication variables
type Jwt struct {
	JWT struct {
		Secret  string        `json:"secret"`
		Expires time.Duration `json:"expired"`
	}
}
