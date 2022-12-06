package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
)

type UnitService interface {
	AddPropertyUnit(*model.PropertyUnitInput) (*model.PropertyUnit, error)
	AddUnitBedrooms([]*model.UnitBedroomInput) ([]*model.Bedroom, error)
	ServiceName() string
}
