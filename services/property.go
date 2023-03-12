package services

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	log "github.com/sirupsen/logrus"
)

var (
	ctx context.Context = context.Background()
)

// PropertyServices - represents property service
type PropertyServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// _ - PropertyServices{} implements PropertyService
var _ interfaces.PropertyService = &PropertyServices{}

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
	creator, err := strconv.ParseInt(property.CreatedBy, 10, 64)
	if err != nil {
		return nil, err
	}
	insertedProperty, err := p.queries.CreateProperty(ctx, sqlStore.CreatePropertyParams{
		Name:       property.Name,
		Town:       property.Town,
		Type:       property.Type,
		MinPrice:   int32(property.MinPrice),
		MaxPrice:   int32(property.MaxPrice),
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
		Type:       insertedProperty.Type,
		MinPrice:   int(insertedProperty.MinPrice),
		MaxPrice:   int(insertedProperty.MaxPrice),
		CreatedBy:  strconv.FormatInt(creator, 10),
		CreatedAt:  &insertedProperty.CreatedAt,
		UpdatedAt:  &insertedProperty.UpdatedAt,
	}, nil
}

// GetProperty - return existing property given property id
func (p *PropertyServices) GetProperty(id string) (*model.Property, error) {
	propertyId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	foundProperty, err := p.queries.GetProperty(ctx, int64(propertyId))
	if err == sql.ErrNoRows {
		// TODO just return empty property list
		return nil, errors.New("Property does not exist")
	}
	return &model.Property{
		ID:         strconv.FormatInt(foundProperty.ID, 10),
		Name:       foundProperty.Name,
		Town:       foundProperty.Town,
		Type:       foundProperty.Type,
		MinPrice:   int(foundProperty.MinPrice),
		MaxPrice:   int(foundProperty.MaxPrice),
		PostalCode: foundProperty.PostalCode,
		CreatedBy:  strconv.FormatInt(foundProperty.CreatedBy, 10),
		CreatedAt:  &foundProperty.CreatedAt,
		UpdatedAt:  &foundProperty.UpdatedAt,
	}, nil
}

// FindByTown - find property(s) in a given town
func (p *PropertyServices) FindByTown(town string) ([]*model.Property, error) {
	return make([]*model.Property, 0), nil
}

// FindByPostalCode - find property(s) in a given postal
func (p *PropertyServices) FindByPostalCode(postalCode string) ([]*model.Property, error) {
	return make([]*model.Property, 0), nil
}

// PropertiesCreatedBy - get property(s) created by user
func (p *PropertyServices) PropertiesCreatedBy(createdBy string) ([]*model.Property, error) {
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
			Type:       item.Type,
			MinPrice:   int(item.MinPrice),
			MaxPrice:   int(item.MaxPrice),
			PostalCode: item.PostalCode,
			CreatedAt:  &item.CreatedAt,
			UpdatedAt:  &item.UpdatedAt,
		}
		userProperties = append(userProperties, property)
	}

	return userProperties, nil
}

// GetPropertyUnits - get property units
func (p *PropertyServices) GetPropertyUnits(propertyId string) ([]*model.PropertyUnit, error) {
	var units []*model.PropertyUnit

	id, err := strconv.ParseInt(propertyId, 10, 64)
	if err != nil {
		return nil, err
	}

	foundUnits, err := p.queries.GetPropertyUnits(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, foundUnit := range foundUnits {
		unit := &model.PropertyUnit{
			ID:         strconv.FormatInt(foundUnit.ID, 10),
			PropertyID: strconv.FormatInt(foundUnit.PropertyID, 10),
			CreatedAt:  &foundUnit.CreatedAt,
			Bathrooms:  int(foundUnit.Bathrooms),
			UpdatedAt:  &foundUnit.UpdatedAt,
		}
		units = append(units, unit)
	}

	return units, nil
}
