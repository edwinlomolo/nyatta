package config

type DatabaseConfig struct {
	RDBMS RDBMS
}

// RDBMS - relational databases variables
type RDBMS struct {
	Uri    string
	Postal struct {
		Uri string
	}
	Env struct {
		Driver string
	}
}
