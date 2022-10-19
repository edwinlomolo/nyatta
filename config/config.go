package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// env - load environment variables
func env() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.WithError(err)
	}
}

// Configuration - load server and db variables
type Configuration struct {
	Database  DatabaseConfig
	JwtConfig Jwt
}

var configAll *Configuration

// Config - load all configurations
func _Config() *Configuration {
	var configuration Configuration

	configuration.Database = database()
	configuration.JwtConfig = jsonWebToken()

	configAll = &configuration

	return configAll
}

// database - all db variables
func database() DatabaseConfig {
	var databaseConfig DatabaseConfig

	// Load environment variables
	env()

	databaseConfig.RDBMS = databaseRDBMS().RDBMS

	return databaseConfig
}

// databaseRDBMS - all relational databases
func databaseRDBMS() DatabaseConfig {
	var databaseConfig DatabaseConfig

	// Load environment variables
	env()

	// Env
	databaseConfig.RDBMS.Env.Driver = os.Getenv("DBDRIVER")
	databaseConfig.RDBMS.Env.Host = os.Getenv("DBHOST")
	databaseConfig.RDBMS.Env.Port = os.Getenv("DBPORT")
	// Access
	databaseConfig.RDBMS.Access.DbName = os.Getenv("DBNAME")
	databaseConfig.RDBMS.Access.User = os.Getenv("DBUSER")
	databaseConfig.RDBMS.Access.Pass = os.Getenv("DBPASS")
	// SSL
	databaseConfig.RDBMS.Ssl.SslMode = os.Getenv("DBSSLMODE")

	return databaseConfig
}

// jsonWebToken - all jwt auth variables
func jsonWebToken() Jwt {
	var jwt Jwt

	// Load env variables
	env()

	jwt.JWT.Expires = os.Getenv("JWTEXPIRE")
	jwt.JWT.Secret = os.Getenv("JWTSECRET")

	return jwt
}

type Config struct {
	// DB
	DBPort     string
	DBUser     string
	DBHost     string
	DBPassword string
	Port       string
	TestDBName string
	DevDBName  string
	ProdDBName string
	SslMode    string

	// Env
	Env string

	// JWT
	JWTSecret     string
	JWTExpiration time.Duration
}

func LoadConfig(path string) (cfg *Config, err error) {
	config := viper.New()
	config.AddConfigPath(path)
	config.SetConfigName(".env")
	config.SetConfigType("env")
	config.AutomaticEnv()
	err = config.ReadInConfig()

	if err != nil {
		// TODO: reuse internal service logger
		log.Fatalf("Error reading env config: %s\n", err)
	}

	cfgs := &Config{
		DBPort:        config.Get("DBPort").(string),
		DBUser:        config.Get("DBUser").(string),
		DBHost:        config.Get("DBHost").(string),
		DBPassword:    config.Get("DBPassword").(string),
		Port:          config.Get("Port").(string),
		SslMode:       config.Get("SslMode").(string),
		Env:           config.Get("Env").(string),
		TestDBName:    config.Get("TestDBName").(string),
		DevDBName:     config.Get("DevDBName").(string),
		ProdDBName:    config.Get("ProdDBName").(string),
		JWTSecret:     config.Get("JWTSecret").(string),
		JWTExpiration: config.GetDuration("JWTExpiration"),
	}
	return cfgs, nil
}
