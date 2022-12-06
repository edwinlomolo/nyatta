package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
)

type TenancyService interface {
	AddUnitTenancy(*model.TenancyInput) (*model.Tenant, error)
	ServiceName() string
}
