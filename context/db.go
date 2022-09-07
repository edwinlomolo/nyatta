package nyatta_context

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

func OpenDB(config *Config, logger *zap.SugaredLogger) (*gorm.DB, error) {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.SslMode)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	logger.Info("Database is connected")
	return db, nil
}
