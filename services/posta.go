package services

import (
	"database/sql"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	log "github.com/sirupsen/logrus"
)

type PostaServices struct {
	Store *sql.DB
}

// _ - PostaServices implement Posta interface
var _ interfaces.Posta = &PostaServices{}

// NewPostaServices - return new posta instance
func NewPostaService() *PostaServices {
	store, err := newPostalStorage()
	if err != nil {
		log.Errorf("%s: %s", config.DatabaseError, err.Error())
	}

	return &PostaServices{Store: store}
}

// newPostalStorage - connects to postal db
func newPostalStorage() (*sql.DB, error) {
	// Connect to postal db
	var dbClient *sql.DB
	dbConfig := config.GetConfig().Database.RDBMS
	//dbUri := fmt.Sprintf("postgres://%s:%s@%s:%s/local", dbConfig.Access.User, dbConfig.Access.Pass, dbConfig.Env.Host, dbConfig.Env.Port)

	// get database instance
	db, err := sql.Open(dbConfig.Env.Driver, dbConfig.Uri)
	if err != nil {
		log.Errorf("%s:%s", config.DatabaseError, err.Error())
	}

	if err := db.Ping(); err == nil {
		dbClient = db
		log.Info("Postal database connected")
	}

	return dbClient, nil
}

// ServiceName - returns service name
func (p PostaServices) ServiceName() string {
	return "PostaServices"
}

// GetTowns - return list of towns
func (p PostaServices) GetTowns() ([]*model.Town, error) {
	var towns []*model.Town
	query := `SELECT id, town, postal_code FROM postal_towns ORDER BY town;`
	rows, err := p.Store.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", config.DatabaseError, err)
	}
	defer rows.Close()

	// Read rows values
	for rows.Next() {
		var id string
		var town string
		var postal_code string

		err = rows.Scan(&id, &town, &postal_code)
		if err != nil {
			return nil, fmt.Errorf("%s:%v", config.DatabaseError, err)
		}
		towns = append(towns, &model.Town{ID: id, Town: town, PostalCode: postal_code})
	}

	// rows.Err will report last error encoutered by rows.Scan
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s:%v", config.DatabaseError, err)
	}

	return towns, nil
}

// SearchTown - get town details
func (p PostaServices) SearchTown(town string) ([]*model.Town, error) {
	var towns []*model.Town
	query := `SELECT id, town, postal_code FROM postal_towns WHERE town ~* $1`
	rows, err := p.Store.Query(query, town)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", config.DatabaseError, err)
	}
	defer rows.Close()

	// Read rows values
	for rows.Next() {
		var id string
		var town string
		var postal_code string

		err = rows.Scan(&id, &town, &postal_code)
		if err != nil {
			return nil, fmt.Errorf("%s:%v", config.DatabaseError, err)
		}
		towns = append(towns, &model.Town{ID: id, Town: town, PostalCode: postal_code})
	}

	// rows.Err will report last error encountered by rows.Scan
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s:%v", config.DatabaseError, err)
	}

	return towns, nil
}
