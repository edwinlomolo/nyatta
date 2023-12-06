package services

import (
	"database/sql"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/sirupsen/logrus"
)

// PostaService - represent posta services
type PostaService interface {
	ServiceName() string
	GetTowns() ([]*model.Town, error)
	SearchTown(town string) ([]*model.Town, error)
}

type postaClient struct {
	Store *sql.DB
	log   *logrus.Logger
}

// NewpostaClient - return new posta instance
func NewPostaService(logger *logrus.Logger) PostaService {
	store, err := newPostalStorage(logger)
	if err != nil {
		logger.Errorf("%s: %s", config.DatabaseError, err.Error())
	}

	return &postaClient{Store: store, log: logger}
}

// newPostalStorage - connects to postal db
func newPostalStorage(logger *logrus.Logger) (*sql.DB, error) {
	// Connect to postal db
	var dbClient *sql.DB
	dbConfig := config.GetConfig().Database.RDBMS

	// get database instance
	db, err := sql.Open(dbConfig.Env.Driver, dbConfig.Postal.Uri)
	if err != nil {
		logger.Errorf("%s:%s", config.DatabaseError, err.Error())
	}

	if err := db.Ping(); err == nil {
		dbClient = db
		logger.Infoln("Postal database connected")
	}

	return dbClient, nil
}

// ServiceName - returns service name
func (p *postaClient) ServiceName() string {
	return "postaClient"
}

// GetTowns - return list of towns
func (p *postaClient) GetTowns() ([]*model.Town, error) {
	var towns []*model.Town
	query := `SELECT id, town, postal_code FROM postal_towns ORDER BY town;`
	rows, err := p.Store.Query(query)
	if err != nil {
		p.log.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}
	defer rows.Close()

	// Read rows values
	for rows.Next() {
		var id string
		var town string
		var postal_code string

		err = rows.Scan(&id, &town, &postal_code)
		if err != nil {
			p.log.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
		towns = append(towns, &model.Town{ID: id, Town: town, PostalCode: postal_code})
	}

	// rows.Err will report last error encoutered by rows.Scan
	if err := rows.Err(); err != nil {
		p.log.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	return towns, nil
}

// SearchTown - get town details
func (p *postaClient) SearchTown(town string) ([]*model.Town, error) {
	var towns []*model.Town
	query := `SELECT id, town, postal_code FROM postal_towns WHERE town ~* $1`
	rows, err := p.Store.Query(query, town)
	if err != nil {
		p.log.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}
	defer rows.Close()

	// Read rows values
	for rows.Next() {
		var id string
		var town string
		var postal_code string

		err = rows.Scan(&id, &town, &postal_code)
		if err != nil {
			p.log.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
		towns = append(towns, &model.Town{ID: id, Town: town, PostalCode: postal_code})
	}

	// rows.Err will report last error encountered by rows.Scan
	if err := rows.Err(); err != nil {
		p.log.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	return towns, nil
}
