package config

type DatabaseConfig struct {
	RDBMS RDBMS `json:"rdbms"`
}

// RDBMS - relational databases variables
type RDBMS struct {
	Uri    string `json:"uri"`
	Postal struct {
		Uri string `json:"uri"`
	}
	Env struct {
		Driver string `json:"driver"`
	}
}
