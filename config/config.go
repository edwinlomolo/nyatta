package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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
	Server    ServerConfig
}

var configAll *Configuration

// Config - load all configurations
func LoadConfig() *Configuration {
	var configuration Configuration

	configuration.Database = database()
	configuration.JwtConfig = jsonWebToken()
	configuration.Server = server()

	configAll = &configuration

	return configAll
}

// GetConfig - get all configurations variables
func GetConfig() *Configuration {
	return configAll
}

// server - all server config variables
func server() ServerConfig {
	var serverConfig ServerConfig

	// Load environment variables
	env()

	serverConfig.ServerPort = os.Getenv("SERVERPORT")
	serverConfig.ServerEnv = os.Getenv("SERVERENV")

	return serverConfig
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
	databaseConfig.RDBMS.Env.SslMode = os.Getenv("SSL_MODE")
	// Access
	databaseConfig.RDBMS.Access.DbName = os.Getenv("DBNAME")
	databaseConfig.RDBMS.Access.User = os.Getenv("DBUSER")
	databaseConfig.RDBMS.Access.Pass = os.Getenv("DBPASS")

	return databaseConfig
}

// jsonWebToken - all jwt auth variables
func jsonWebToken() Jwt {
	var jwt Jwt

	// Load env variables
	env()

	duration, err := time.ParseDuration(os.Getenv("JWTEXPIRE"))
	if err != nil {
		log.Errorf("panic: jwt duration: %v", err)
	}

	jwt.JWT.Expires = duration
	jwt.JWT.Secret = os.Getenv("JWTSECRET")

	return jwt
}

// IsPrototypeEnv - is environment development?
func IsPrototypeEnv() bool {
	// Load environment variables
	env()

	serverEnv := os.Getenv("SERVERENV")

	return (serverEnv == "development" || serverEnv == "test")
}
