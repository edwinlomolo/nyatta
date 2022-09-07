package nyatta_context

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBName     string
	DBPort     string
	DBUser     string
	DBHost     string
	DBPassword string
	Port       string
	Env        string

	SslMode string
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

	config.Unmarshal(&cfg)
	return
}
