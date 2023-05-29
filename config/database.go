package config

type DatabaseConfig struct {
	RDBMS RDBMS
}

// RDBMS - relational databases variables
type RDBMS struct {
	Uri string
	Env struct {
		Driver  string
		Host    string
		Port    string
		SslMode string
	}
	Access struct {
		DbName string
		User   string
		Pass   string
	}
}
