package services

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	log "github.com/sirupsen/logrus"
)

type PropertyService interface {
	ServiceName() string
	CreateProperty(*model.NewProperty) (*sqlStore.Property, error)
	GetProperty(id string) (*sqlStore.Property, error)
	FindByTown(town string) ([]*model.Property, error)
	FindByPostalCode(postalCode string) ([]*model.Property, error)
	// AddAmenity(*model.AmenityInput) (*model.Amenity, error)
}

// PropertyServices - represents property service
type PropertyServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// _ - PropertyServices{} implements PropertyService
var _ PropertyService = &PropertyServices{}

// NewPropertyService - factory for property services
func NewPropertyService(queries *sqlStore.Queries, logger *log.Logger) *PropertyServices {
	return &PropertyServices{queries: queries, logger: logger}
}

// ServiceName - return service name
func (p PropertyServices) ServiceName() string {
	return "PropertyServices"
}

// CreateProperty - create new property
func (p *PropertyServices) CreateProperty(property *model.NewProperty) (*sqlStore.Property, error) {
	ctx := context.Background()

	creator, err := strconv.ParseInt(property.CreatedBy, 10, 64)
	if err != nil {
		return nil, err
	}
	insertedProperty, err := p.queries.CreateProperty(ctx, sqlStore.CreatePropertyParams{
		Name:       property.Name,
		Town:       property.Town,
		PostalCode: property.PostalCode,
		CreatedBy:  sql.NullInt64{Int64: creator, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &sqlStore.Property{
		ID:         insertedProperty.ID,
		Name:       insertedProperty.Name,
		Town:       insertedProperty.Town,
		PostalCode: insertedProperty.PostalCode,
		CreatedBy:  insertedProperty.CreatedBy,
	}, nil
}

// GetProperty - return existing property given property id
func (p *PropertyServices) GetProperty(id string) (*sqlStore.Property, error) {
	ctx := context.Background()
	propertyId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	foundProperty, err := p.queries.GetProperty(ctx, int64(propertyId))
	if err == sql.ErrNoRows {
		return nil, errors.New("Property does not exist")
	}
	return &sqlStore.Property{
		ID:         foundProperty.ID,
		Name:       foundProperty.Name,
		Town:       foundProperty.Town,
		PostalCode: foundProperty.PostalCode,
	}, nil
}

// TODO
// FindByTown - find property(s) in a given town
func (p *PropertyServices) FindByTown(town string) ([]*model.Property, error) {
	return make([]*model.Property, 0), nil
}

// TODO
// FindByPostalCode - find property(s) in a given postal
func (p *PropertyServices) FindByPostalCode(postalCode string) ([]*model.Property, error) {
	return make([]*model.Property, 0), nil
}

// AddAmenity - add property amenity(s)
/*
func (p *PropertyServices) AddAmenity(amenity *model.AmenityInput) (*model.Amenity, error) {
	newAmenity := &model.Amenity{
		Name:       amenity.Name,
		Provider:   amenity.Provider,
		PropertyID: amenity.PropertyID,
	}
	err := p.store.Create(&newAmenity).Error
	if err != nil {
		return nil, err
	}

	return newAmenity, nil
}
*/
