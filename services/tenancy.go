package services

import (
	"context"
	"database/sql"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// TenancyServices - represent tenancy services
type TenancyService interface {
	ServiceName() string
	AddUnitTenancy(ctx context.Context, input *model.TenancyInput) (*model.Tenant, error)
	GetUnitTenancy(ctx context.Context, unitID uuid.UUID) ([]*model.Tenant, error)
	GetCurrentTenant(ctx context.Context, unitID uuid.UUID) (*model.Tenant, error)
	GetUserTenancy(ctx context.Context, userID uuid.UUID) ([]*model.Tenant, error)
}

// NewTenancyService - factory for tenancy services
func NewTenancyService(queries *sqlStore.Queries, logger *log.Logger) TenancyService {
	return &tenancyClient{queries: queries, logger: logger}
}

type tenancyClient struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// AddUnitTenancy - add tenancy to property unit
func (t *tenancyClient) AddUnitTenancy(ctx context.Context, input *model.TenancyInput) (*model.Tenant, error) {
	user, err := ctx.Value("userService").(UserService).FindUserByPhone(ctx, input.Phone)
	if err != nil {
		t.logger.Errorf("%s:%v", t.ServiceName(), err)
		return nil, err
	}

	unit, err := ctx.Value("unitService").(UnitService).GetUnit(ctx, input.UnitID)
	if err != nil {
		t.logger.Errorf("%s:%v", t.ServiceName(), err)
		return nil, err
	}

	insertedTenant, err := t.queries.CreateTenant(ctx, sqlStore.CreateTenantParams{
		PropertyID: unit.PropertyID,
		UnitID:     input.UnitID,
		UserID:     user.ID,
		StartDate:  input.StartDate,
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
func (t *tenancyClient) GetUnitTenancy(ctx context.Context, unitId uuid.UUID) ([]*model.Tenant, error) {
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
func (t *tenancyClient) GetCurrentTenant(ctx context.Context, unitID uuid.UUID) (*model.Tenant, error) {
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
func (t *tenancyClient) ServiceName() string {
	return "tenancyClient"
}

// GetUserTenancy - grab user tenancy history
func (t *tenancyClient) GetUserTenancy(ctx context.Context, userID uuid.UUID) ([]*model.Tenant, error) {
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
