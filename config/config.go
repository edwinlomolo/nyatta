package config

import (
	"encoding/base64"
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
		log.Errorf("Error loading env vars: %v", err)
	}
}

// Configuration - load server and db variables
type Configuration struct {
	Database     DatabaseConfig `json:"database"`
	JwtConfig    Jwt            `json:"jwt"`
	Server       ServerConfig   `json:"server"`
	Aws          AwsConfig      `json:"aws"`
	Twilio       TwilioConfig   `json:"twilio"`
	SentryConfig SentryConfig   `json:"sentry"`
	Email        EmailConfig    `json:"email"`
	Mpesa        MpesaConfig    `json:"mpesa"`
	Paystack     Paystack       `json:"paystack"`
	EquityBank   EquityBank     `json:"equityBank"`
}

var configAll *Configuration

// Config - load all configurations
func LoadConfig() *Configuration {
	var configuration Configuration

	configuration.Database = database()
	configuration.JwtConfig = jsonWebToken()
	configuration.Server = server()
	configuration.Aws = awsConfig()
	configuration.Twilio = twilioConfig()
	configuration.SentryConfig = sentryConfig()
	configuration.Email = emailConfig()
	configuration.Mpesa = mpesaConfig()
	configuration.Paystack = paystackConfig()
	configuration.EquityBank = equityBankConfig()

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
	aws.S3.Buckets.Media = os.Getenv("S3_BUCKET")

	return aws
}

// twilioConfig - setup twilio config
func twilioConfig() TwilioConfig {
	var twilio TwilioConfig

	// Load env variables
	env()

	twilio.Sid = os.Getenv("TWILIO_ACCOUNT_SID")
	twilio.AuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	twilio.VerifySid = os.Getenv("TWILIO_VERIFICATION_SID")

	return twilio
}

// ForcePostgresMigration - force postgres migration
func ForcePostgresMigration() bool {

	// Load env variables
	env()

	forceMigration := os.Getenv("FORCE_MIGRATION")

	return forceMigration == "true"
}

// sentryConfig - setup sentry config
func sentryConfig() SentryConfig {
	var sentryConfig SentryConfig

	// Load env variables
	env()

	sentryConfig.Dsn = os.Getenv("SENTRY_DSN")

	return sentryConfig
}

// emailConfig - setup email config
func emailConfig() EmailConfig {
	var emailConfig EmailConfig

	// Load env variables
	env()

	emailConfig.From = os.Getenv("EMAIL_FROM")
	emailConfig.Apikey = os.Getenv("RESEND_API_KEY")

	return emailConfig
}

// mpesaConfig - setup mpesa config
func mpesaConfig() MpesaConfig {
	var mpesaConfig MpesaConfig

	// Load env variables
	env()

	mpesaConfig.ConsumerKey = os.Getenv("MPESA_CONSUMER_KEY")
	mpesaConfig.ConsumerSecret = os.Getenv("MPESA_CONSUMER_SECRET")
	mpesaConfig.BaseApi = os.Getenv("MPESA_BASE_API")
	mpesaConfig.PassKey = os.Getenv("MPESA_PASS_KEY")

	return mpesaConfig
}

// paystackConfig - setup paystack config
func paystackConfig() Paystack {
	var paystackConfig Paystack

	// Load env
	env()

	paystackConfig.SecretKey = os.Getenv("PAYSTACK_SECRET_KEY")
	paystackConfig.BaseApi = os.Getenv("PAYSTACK_BASE_API")

	return paystackConfig
}

// equityBankConfig - setup equity bank config
func equityBankConfig() EquityBank {
	var equityBankConfig EquityBank

	// Load env
	env()

	equityBankConfig.MerchantCode = os.Getenv("EQUITY_MERCHANT_CODE")
	equityBankConfig.ConsumerSecret = os.Getenv("EQUITY_CONSUMER_SECRET")
	equityBankConfig.ApiKey = os.Getenv("EQUITY_API_KEY")
	privateKey, err := base64.StdEncoding.DecodeString(os.Getenv("EQUITY_PRIVATE_KEY"))
	if err != nil {
		log.Fatalln("Error reading private key env")
	}
	equityBankConfig.PrivateKey = string(privateKey)
	equityBankConfig.BaseApi = os.Getenv("EQUITY_BASE_API")

	return equityBankConfig
}
