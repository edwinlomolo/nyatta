package interfaces

import "github.com/3dw1nM0535/nyatta/graph/model"

type PropertyService interface {
	ServiceName() string
	CreateProperty(*model.NewProperty) (*model.Property, error)
	GetProperty(id string) (*model.Property, error)
	FindByTown(town string) ([]*model.Property, error)
	FindByPostalCode(postalCode string) ([]*model.Property, error)
	PropertiesCreatedBy(createdBy string) ([]*model.Property, error)
}