package services

import (
	"errors"

	"github.com/3dw1nM0535/nyatta/graph/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PropertyService interface {
	ServiceName() string
	CreateProperty(*model.NewProperty) (*model.Property, error)
	GetProperty(id string) (*model.Property, error)
	FindByTown(town string) ([]*model.Property, error)
	FindByPostalCode(postalCode string) ([]*model.Property, error)
	AddAmenity(*model.AmenityInput) (*model.Amenity, error)
}

// PropertyServices - represents property service
type PropertyServices struct {
	store  *gorm.DB
	logger *log.Logger
}

// _ - PropertyServices{} implements PropertyService
var _ PropertyService = &PropertyServices{}

// NewPropertyService - factory for property services
func NewPropertyService(store *gorm.DB, logger *log.Logger) *PropertyServices {
	return &PropertyServices{store: store, logger: logger}
}

// ServiceName - return service name
func (p PropertyServices) ServiceName() string {
	return "PropertyServices"
}

// CreateProperty - create new property
func (p *PropertyServices) CreateProperty(property *model.NewProperty) (*model.Property, error) {
	newProperty := &model.Property{
		Name:       property.Name,
		Town:       property.Town,
		PostalCode: property.PostalCode,
		CreatedBy:  property.CreatedBy,
	}

	err := p.store.Create(&newProperty).Error
	if err != nil {
		return nil, err
	}

	return newProperty, nil
}

// GetProperty - return existing property given property id
func (p *PropertyServices) GetProperty(id string) (*model.Property, error) {
	var foundProperty *model.Property
	err := p.store.Where("id = ?", id).Preload("Amenities").Find(&foundProperty).Error
	if err != nil {
		return nil, err
	}

	if foundProperty.ID != id {
		return nil, errors.New("Property does not exist")
	}

	return foundProperty, nil
}

// FindByTown - find property(s) in a given town
func (p *PropertyServices) FindByTown(town string) ([]*model.Property, error) {
	var foundProperties []*model.Property

	err := p.store.Where("town = ?", town).Preload("Amenities").Find(&foundProperties).Error
	if err != nil {
		return nil, err
	}

	return foundProperties, nil
}

// FindByPostalCode - find property(s) in a given postal
func (p *PropertyServices) FindByPostalCode(postalCode string) ([]*model.Property, error) {
	var foundProperties []*model.Property

	err := p.store.Where("postal_code = ?", postalCode).Preload("Amenities").Find(&foundProperties).Error
	if err != nil {
		return nil, err
	}

	return foundProperties, nil
}

// AddAmenity - add property amenity(s)
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
