package interfaces

import "github.com/3dw1nM0535/nyatta/graph/model"

type PropertyService interface {
	ServiceName() string
	CreateProperty(*model.NewProperty) (*model.Property, error)
	GetProperty(id string) (*model.Property, error)
	PropertiesCreatedBy(createdBy string) ([]*model.Property, error)
	GetPropertyUnits(propertyId string) ([]*model.PropertyUnit, error)
	CaretakerPhoneVerification(*model.CaretakerVerificationInput) (*model.Status, error)
	ListingOverview(propertyId string) (*model.ListingOverview, error)
}
