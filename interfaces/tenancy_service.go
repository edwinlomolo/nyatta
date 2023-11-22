package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
)

type TenancyService interface {
	AddUnitTenancy(*model.TenancyInput) (*model.Tenant, error)
	ServiceName() string
	GetCurrentTenant(unitID uuid.UUID) (*model.Tenant, error)
	GetUnitTenancy(unitId uuid.UUID) ([]*model.Tenant, error)
}
