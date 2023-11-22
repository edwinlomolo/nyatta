package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
)

type UnitService interface {
	AddPropertyUnit(ctx context.Context, unit *model.PropertyUnitInput) (*model.PropertyUnit, error)
	GetUnitBedrooms(ctx context.Context, unitId uuid.UUID) ([]*model.Bedroom, error)
	GetUnitImages(ctx context.Context, id uuid.UUID) ([]*model.AnyUpload, error)
	ServiceName() string
}
