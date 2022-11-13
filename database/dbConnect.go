package database

import (
	"errors"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DatabaseError = errors.New("DatabaseError")
)

var dbClient *gorm.DB

// InitDB - setup db and return connection instance/error
func InitDB() (*gorm.DB, error) {
	configureDB := config.GetConfig().Database.RDBMS

	host := configureDB.Env.Host
	port := configureDB.Env.Port
	user := configureDB.Access.User
	pass := configureDB.Access.Pass
	name := configureDB.Access.DbName
	ssl_mode := configureDB.Ssl.SslMode

	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl_mode)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(1)),
	})
	if err != nil {
		return nil, err
	}
	dbClient = db
	log.Info("Database is connected")
	if err := dropAllTables(dbClient); err != nil {
		return nil, err
	}

	if err := startMigration(dbClient); err != nil {
		return nil, err
	}
	return dbClient, nil
}

// GetDB - get database client
func GetDB() *gorm.DB {
	return dbClient
}

// dropAllTables - cleanup database tables
func dropAllTables(db *gorm.DB) error {
	if err := db.Migrator().DropTable(
		&model.User{},
		&model.Property{},
		&model.Amenity{},
	); err != nil {
		log.WithError(err)
		return err
	}
	log.Info("Database tables deleted")
	return nil
}

// startMigration - setup database tables/columns and any missing indexes
func startMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
		&model.Property{},
		&model.Amenity{},
	); err != nil {
		log.WithError(err)
		return err
	}
	log.Info("Database tables migrated")
	return nil
}
