package services

import (
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type TenancyServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// _ - TenancyServices{} implements TenancyService interface
var _ interfaces.TenancyService = &TenancyServices{}

// NewTenancyService - factory for tenancy services
func NewTenancyService(queries *sqlStore.Queries, logger *log.Logger) *TenancyServices {
	return &TenancyServices{queries, logger}
}

// AddUnitTenancy - add tenancy to property unit
func (u *TenancyServices) AddUnitTenancy(input *model.TenancyInput) (*model.Tenant, error) {
	insertedTenant, err := u.queries.CreateTenant(ctx, sqlStore.CreateTenantParams{
		PropertyUnitID: uuid.NullUUID{UUID: input.PropertyUnitID, Valid: true},
		StartDate:      input.StartDate,
	})
	if err != nil {
		u.logger.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}
	return &model.Tenant{
		ID:             insertedTenant.ID,
		StartDate:      insertedTenant.StartDate,
		EndDate:        &insertedTenant.EndDate.Time,
		PropertyUnitID: input.PropertyUnitID,
		CreatedAt:      &insertedTenant.CreatedAt,
		UpdatedAt:      &insertedTenant.UpdatedAt,
	}, nil
}

// ServiceName - return service name
func (t TenancyServices) ServiceName() string {
	return "TenancyServices"
}
