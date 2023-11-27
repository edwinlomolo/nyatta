package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
)

type AmenityService interface {
	AddAmenity(ctx context.Context, unitID uuid.UUID, input *model.UnitAmenityInput) (*model.Amenity, error)
	GetUnitAmenities(ctx context.Context, unitID uuid.UUID) ([]*model.Amenity, error)
}
