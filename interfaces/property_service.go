package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
)

type PropertyService interface {
	ServiceName() string
	CreateProperty(property *model.NewProperty, isLandlord bool, phone string, createdBy uuid.UUID) (*model.Property, error)
	GetProperty(id uuid.UUID) (*model.Property, error)
	GetPropertyThumbnail(id uuid.UUID) (*model.AnyUpload, error)
	PropertiesCreatedBy(createdBy uuid.UUID) ([]*model.Property, error)
	GetPropertyUnits(propertyId uuid.UUID) ([]*model.PropertyUnit, error)
	CaretakerPhoneVerification(*model.CaretakerVerificationInput) (*model.Status, error)
	ListingOverview(propertyId uuid.UUID) (*model.ListingOverview, error)
	CreatePropertyCaretaker(propertyId uuid.UUID) (*model.Caretaker, error)
	GetPropertyCaretaker(caretakerId uuid.UUID) (*model.Caretaker, error)
}
