package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
)

type TenancyService interface {
	AddUnitTenancy(ctx context.Context, tenant *model.TenancyInput) (*model.Tenant, error)
	ServiceName() string
	GetCurrentTenant(ctx context.Context, unitID uuid.UUID) (*model.Tenant, error)
	GetUnitTenancy(ctx context.Context, unitId uuid.UUID) ([]*model.Tenant, error)
}
