package nyatta_context

import (
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

func OpenDB(config *Config, logger *zap.SugaredLogger) (*gorm.DB, error) {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DevDBName, config.SslMode)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	logger.Info("Database is connected")
	db.Migrator().DropTable(
		&model.User{},
		&model.Property{},
	)
	if err := AutoMigrate(db); err != nil {
		logger.Errorf("%s: %v", DatabaseError, err)
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Property{},
	)
}
