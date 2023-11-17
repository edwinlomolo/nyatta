package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

var (
	DatabaseError = errors.New("DatabaseError")
)

var dbClient *sql.DB

// InitDB - setup db and return connection instance/error
func InitDB(migrationUrl string) (*sql.DB, error) {
	configureDB := config.GetConfig().Database.RDBMS

	uri := configureDB.Uri
	driver := configureDB.Env.Driver

	db, err := sql.Open(driver, uri)
	if err != nil {
		log.Errorf("%s:%s", config.DatabaseError, err.Error())
		return nil, err
	}

	status := "up"
	if err := db.Ping(); err != nil {
		status = "down"
		log.Errorf("%s:%s", config.DatabaseError, err.Error())
	}

	dbClient = db
	if dbClient == nil {
		log.Errorf("%s:%s", config.DatabaseError, "nil db instance")
	}
	log.Infof("Database is %s", status)

	// Setup database schema
	if err := runDbMigration(dbClient, migrationUrl); err == nil {
		log.Infoln("Database migration applied")
	}

	return dbClient, nil
}

// GetDB - get database client
func GetDB() *sql.DB {
	return dbClient
}

// runDbMigration - setup database tables
func runDbMigration(db *sql.DB, migrationUrl string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Errorf("%s: %s", config.MigrationDriverErr, err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationUrl), "postgres", driver)
	if err != nil {
		log.Errorf("%s: %s", config.MigrationInstanceErr, err)
		return err
	}

	// Apply migration(s)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Errorf("%s: %s", config.MigrationUpErr, err)
		return err
	}

	if config.IsPrototypeEnv() && config.ForcePostgresMigration() {
		// Drop everything
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Errorf("%s: %s", config.MigrationDownErr, err)
			return err
		}
	}

	/*
		if !config.IsPrototypeEnv() {
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Errorf("%s: %s", config.MigrationErr, err)
				return err
			}
		}
	*/
	return nil
}
