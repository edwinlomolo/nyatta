package config

// JWT - jwt authentication variables
type Jwt struct {
	JWT struct {
		Secret  string
		Expires string
	}
}
