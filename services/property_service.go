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
	CreateProperty(*model.NewProperty) (*model.Property, error)
	GetProperty(id string) (*model.Property, error)
	FindByTown(town string) ([]*model.Property, error)
	FindByPostalCode(postalCode string) ([]*model.Property, error)
	PropertiesCreatedBy(createdBy string) ([]*model.Property, error)
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
func (p *PropertyServices) CreateProperty(property *model.NewProperty) (*model.Property, error) {
	ctx := context.Background()

	creator, err := strconv.ParseInt(property.CreatedBy, 10, 64)
	if err != nil {
		return nil, err
	}
	insertedProperty, err := p.queries.CreateProperty(ctx, sqlStore.CreatePropertyParams{
		Name:       property.Name,
		Town:       property.Town,
		PostalCode: property.PostalCode,
		CreatedBy:  creator,
	})
	if err != nil {
		return nil, err
	}

	return &model.Property{
		ID:         strconv.FormatInt(insertedProperty.ID, 10),
		Name:       insertedProperty.Name,
		Town:       insertedProperty.Town,
		PostalCode: insertedProperty.PostalCode,
		CreatedBy:  strconv.FormatInt(insertedProperty.CreatedBy, 10),
	}, nil
}

// GetProperty - return existing property given property id
func (p *PropertyServices) GetProperty(id string) (*model.Property, error) {
	ctx := context.Background()
	propertyId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	foundProperty, err := p.queries.GetProperty(ctx, int64(propertyId))
	if err == sql.ErrNoRows {
		return nil, errors.New("Property does not exist")
	}
	return &model.Property{
		ID:         strconv.FormatInt(foundProperty.ID, 10),
		Name:       foundProperty.Name,
		Town:       foundProperty.Town,
		PostalCode: foundProperty.PostalCode,
		CreatedAt:  &foundProperty.CreatedAt,
		UpdatedAt:  &foundProperty.UpdatedAt,
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

func (p *PropertyServices) PropertiesCreatedBy(createdBy string) ([]*model.Property, error) {
	ctx := context.Background()
	var userProperties []*model.Property

	// Use int64 id
	creator, err := strconv.ParseInt(createdBy, 10, 64)
	if err != nil {
		return nil, err
	}

	props, err := p.queries.PropertiesCreatedBy(ctx, creator)
	if err == sql.ErrNoRows {
		return nil, errors.New("No properties found")
	}

	for _, item := range props {
		property := &model.Property{
			ID:         strconv.FormatInt(item.ID, 10),
			Name:       item.Name,
			Town:       item.Town,
			PostalCode: item.PostalCode,
			CreatedAt:  &item.CreatedAt,
			UpdatedAt:  &item.UpdatedAt,
		}
		userProperties = append(userProperties, property)
	}

	return userProperties, nil
}
