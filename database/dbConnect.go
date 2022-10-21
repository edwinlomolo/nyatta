package database

import (
	"errors"
	"fmt"

	cfg "github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DatabaseError = errors.New("DatabaseError")
)

func InitDB(config *cfg.RDBMS) (*gorm.DB, error) {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Env.Host, config.Env.Port, config.Access.User, config.Access.Pass, config.Access.DbName, config.Ssl.SslMode)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(1)),
	})
	if err != nil {
		return nil, err
	}
	log.Info("Database is connected")
	db.Migrator().DropTable(
		&model.User{},
		&model.Property{},
		&model.Amenity{},
	)
	if err := AutoMigrate(db); err != nil {
		log.Errorf("%s: %v", DatabaseError, err)
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Property{},
		&model.Amenity{},
	)
}
