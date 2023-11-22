package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
)

type UnitService interface {
	AddPropertyUnit(*model.PropertyUnitInput) (*model.PropertyUnit, error)
	GetUnitBedrooms(unitId uuid.UUID) ([]*model.Bedroom, error)
	GetUnitImages(id uuid.UUID) ([]*model.AnyUpload, error)
	ServiceName() string
}
