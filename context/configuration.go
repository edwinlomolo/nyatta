package nyatta_context

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

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
