package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
)

type PropertyService interface {
	ServiceName() string
	CreateProperty(ctx context.Context, property *model.NewProperty, createdBy uuid.UUID) (*model.Property, error)
	GetProperty(ctx context.Context, id uuid.UUID) (*model.Property, error)
	GetPropertyThumbnail(ctx context.Context, id uuid.UUID) (*model.AnyUpload, error)
	PropertiesCreatedBy(ctx context.Context, createdBy uuid.UUID) ([]*model.Property, error)
	GetUnits(ctx context.Context, propertyId uuid.UUID) ([]*model.Unit, error)
	CaretakerPhoneVerification(context.Context, *model.CaretakerVerificationInput) (*model.Status, error)
	ListingOverview(ctx context.Context, propertyId uuid.UUID) (*model.ListingOverview, error)
	GetPropertyCaretaker(ctx context.Context, caretakerId uuid.UUID) (*model.Caretaker, error)
}
