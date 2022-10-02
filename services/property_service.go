package services

import (
	"errors"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"go.uber.org/zap"
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

type PropertyServices struct {
	store  *gorm.DB
	logger *zap.SugaredLogger
}

func NewPropertyService(store *gorm.DB, logger *zap.SugaredLogger) *PropertyServices {
	return &PropertyServices{store: store, logger: logger}
}

func (p PropertyServices) ServiceName() string {
	return "PropertyServices"
}

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

func (p *PropertyServices) FindByTown(town string) ([]*model.Property, error) {
	var foundProperties []*model.Property

	err := p.store.Where("town = ?", town).Find(&foundProperties).Error
	if err != nil {
		return nil, err
	}

	return foundProperties, nil
}

func (p *PropertyServices) FindByPostalCode(postalCode string) ([]*model.Property, error) {
	var foundProperties []*model.Property

	err := p.store.Where("postal_code = ?", postalCode).Find(&foundProperties).Error
	if err != nil {
		return nil, err
	}

	return foundProperties, nil
}

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
