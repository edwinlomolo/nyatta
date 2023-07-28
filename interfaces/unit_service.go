package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
)

type UnitService interface {
	AddPropertyUnit(*model.PropertyUnitInput) (*model.PropertyUnit, error)
	AddUnitBedrooms([]*model.UnitBedroomInput) ([]*model.Bedroom, error)
	GetUnitBedrooms(unitId string) ([]*model.Bedroom, error)
	GetUnitTenancy(unitId string) ([]*model.Tenant, error)
	AmenityCount(unitId string) (int64, error)
	ServiceName() string
}
