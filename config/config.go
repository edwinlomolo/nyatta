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
	Aws       AwsConfig
}

var configAll *Configuration

// Config - load all configurations
func LoadConfig() *Configuration {
	var configuration Configuration

	configuration.Database = database()
	configuration.JwtConfig = jsonWebToken()
	configuration.Server = server()
	configuration.Aws = awsConfig()

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
	databaseConfig.RDBMS.Postal.Uri = os.Getenv("POSTAL_DATABASE_URI")
	databaseConfig.RDBMS.Uri = os.Getenv("DATABASE_URI")
	databaseConfig.RDBMS.Env.Driver = os.Getenv("DBDRIVER")

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

// awsConfig - setup aws config
func awsConfig() AwsConfig {
	var aws AwsConfig

	// Load env variables
	env()

	aws.AccessKey = os.Getenv("ACCESS_KEY")
	aws.SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY")

	return aws
}
