package services

import (
	"strconv"

	"database/sql"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
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
	propertyId, err := strconv.ParseInt(input.PropertyUnitID, 10, 64)
	if err != nil {
		return nil, err
	}

	insertedTenant, err := u.queries.CreateTenant(ctx, sqlStore.CreateTenantParams{
		PropertyUnitID: propertyId,
		StartDate:      input.StartDate,
		EndDate:        sql.NullTime{Time: *input.EndDate, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &model.Tenant{
		ID:             strconv.FormatInt(insertedTenant.ID, 10),
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
