package services

import (
	"context"
	"database/sql"

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
func (t *TenancyServices) AddUnitTenancy(ctx context.Context, input *model.TenancyInput) (*model.Tenant, error) {
	insertedTenant, err := t.queries.CreateTenant(ctx, sqlStore.CreateTenantParams{
		UnitID:    input.UnitID,
		UserID:    input.UserID,
		StartDate: input.StartDate,
	})
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return nil, err
	}
	return &model.Tenant{
		ID:        insertedTenant.ID,
		StartDate: insertedTenant.StartDate,
		EndDate:   &insertedTenant.EndDate.Time,
		UnitID:    input.UnitID,
		CreatedAt: &insertedTenant.CreatedAt,
		UpdatedAt: &insertedTenant.UpdatedAt,
	}, nil
}

// GetUnitTenancy - return unit tenancy
func (t *TenancyServices) GetUnitTenancy(ctx context.Context, unitId uuid.UUID) ([]*model.Tenant, error) {
	var tenancies []*model.Tenant

	foundTenancies, err := t.queries.GetUnitTenancy(ctx, unitId)
	if err != nil {
		if err == sql.ErrNoRows {
			return tenancies, nil
		} else {
			t.logger.Errorf("%s:%v", t.ServiceName(), err)
			return nil, err
		}
	}

	for _, tenancy := range foundTenancies {
		tenant := &model.Tenant{
			ID:        tenancy.ID,
			StartDate: tenancy.StartDate,
			UserID:    tenancy.UserID,
			UnitID:    tenancy.UnitID,
			EndDate:   &tenancy.EndDate.Time,
			CreatedAt: &tenancy.CreatedAt,
			UpdatedAt: &tenancy.UpdatedAt,
		}
		tenancies = append(tenancies, tenant)
	}

	return tenancies, nil
}

// GetUnitTenant - grab current tenant
func (t *TenancyServices) GetCurrentTenant(ctx context.Context, unitID uuid.UUID) (*model.Tenant, error) {
	foundTenant, err := t.queries.GetCurrentTenant(ctx, unitID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			t.logger.Errorf("%s:%v", t.ServiceName(), err)
			return nil, err
		}
	}

	return &model.Tenant{
		ID:        foundTenant.ID,
		StartDate: foundTenant.StartDate,
		UserID:    foundTenant.UserID,
		UnitID:    foundTenant.UnitID,
		EndDate:   &foundTenant.EndDate.Time,
		CreatedAt: &foundTenant.CreatedAt,
		UpdatedAt: &foundTenant.UpdatedAt,
	}, nil
}

// ServiceName - return service name
func (t TenancyServices) ServiceName() string {
	return "TenancyServices"
}

// GetUserTenancy - grab user tenancy history
func (t *TenancyServices) GetUserTenancy(ctx context.Context, userID uuid.UUID) ([]*model.Tenant, error) {
	var tenancy []*model.Tenant

	foundTenancy, err := t.queries.GetUserTenancy(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return tenancy, nil
		} else {
			t.logger.Errorf("%s:%v", t.ServiceName(), err)
			return nil, err
		}
	}

	for _, tenancyItem := range foundTenancy {
		tenant := &model.Tenant{
			ID:        tenancyItem.ID,
			StartDate: tenancyItem.StartDate,
			EndDate:   &tenancyItem.EndDate.Time,
		}
		tenancy = append(tenancy, tenant)
	}

	return tenancy, nil
}
