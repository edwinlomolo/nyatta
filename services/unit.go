package services

import (
	"errors"
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	log "github.com/sirupsen/logrus"
)

type UnitServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// _ - UnitServices{} implements UnitService interface
var _ interfaces.UnitService = &UnitServices{}

// NewUnitService - factory for UnitServices
func NewUnitService(queries *sqlStore.Queries, logger *log.Logger) *UnitServices {
	return &UnitServices{queries, logger}
}

// AddPropertyUnit - add property unit
func (u *UnitServices) AddPropertyUnit(input *model.PropertyUnitInput) (*model.PropertyUnit, error) {
	if input.Bathrooms == 0 {
		return nil, errors.New("Zero is not a valid value")
	}
	propertyId, err := strconv.ParseInt(input.PropertyID, 10, 64)
	if err != nil {
		return nil, err
	}

	insertedUnit, err := u.queries.CreatePropertyUnit(ctx, sqlStore.CreatePropertyUnitParams{
		PropertyID: propertyId,
		Bathrooms:  int32(input.Bathrooms),
	})
	if err != nil {
		return nil, err
	}

	return &model.PropertyUnit{
		ID:         strconv.FormatInt(insertedUnit.ID, 10),
		Bathrooms:  int(insertedUnit.Bathrooms),
		PropertyID: input.PropertyID,
		CreatedAt:  &insertedUnit.CreatedAt,
		UpdatedAt:  &insertedUnit.UpdatedAt,
	}, nil
}

// AddUnitBedroom - add property unit bedroom(s)
func (u *UnitServices) AddUnitBedrooms(input []*model.UnitBedroomInput) ([]*model.Bedroom, error) {
	var insertedBedrooms []*model.Bedroom
	for _, value := range input {
		if value.BedroomNumber == 0 {
			return nil, errors.New("Zero is not a valid value")
		}
		propertyUnitId, err := strconv.ParseInt(value.PropertyUnitID, 10, 64)
		if err != nil {
			return nil, err
		}
		insertedBedroom, err := u.queries.CreateUnitBedroom(ctx, sqlStore.CreateUnitBedroomParams{
			PropertyUnitID: propertyUnitId,
			BedroomNumber:  int32(value.BedroomNumber),
			EnSuite:        value.EnSuite,
			Master:         value.Master,
		})
		if err != nil {
			return nil, err
		}
		insertedBedrooms = append(insertedBedrooms, &model.Bedroom{
			ID:        strconv.FormatInt(insertedBedroom.ID, 10),
			EnSuite:   insertedBedroom.EnSuite,
			Master:    insertedBedroom.Master,
			CreatedAt: &insertedBedroom.CreatedAt,
			UpdatedAt: &insertedBedroom.UpdatedAt,
		})
	}
	return insertedBedrooms, nil
}

// GetUnitBedrooms - return unit bedrooms
func (u *UnitServices) GetUnitBedrooms(unitId string) ([]*model.Bedroom, error) {
	var bedrooms []*model.Bedroom
	id, err := strconv.ParseInt(unitId, 10, 64)
	if err != nil {
		return nil, err
	}

	foundBedrooms, err := u.queries.GetUnitBedrooms(ctx, id)
	for _, unit := range foundBedrooms {
		bedroom := &model.Bedroom{
			ID:            strconv.FormatInt(unit.ID, 10),
			BedroomNumber: int(unit.BedroomNumber),
			EnSuite:       unit.EnSuite,
			Master:        unit.Master,
			CreatedAt:     &unit.CreatedAt,
			UpdatedAt:     &unit.UpdatedAt,
		}
		bedrooms = append(bedrooms, bedroom)
	}
	return bedrooms, nil
}

// GetUnitTenancy - return unit tenancy
func (u *UnitServices) GetUnitTenancy(unitId string) ([]*model.Tenant, error) {
	var tenancies []*model.Tenant

	id, err := strconv.ParseInt(unitId, 10, 64)
	if err != nil {
		return nil, err
	}

	foundTenancies, err := u.queries.GetUnitTenancy(ctx, id)
	for _, tenancy := range foundTenancies {
		tenant := &model.Tenant{
			ID:        strconv.FormatInt(tenancy.ID, 10),
			StartDate: tenancy.StartDate,
			EndDate:   &tenancy.EndDate.Time,
			CreatedAt: &tenancy.CreatedAt,
			UpdatedAt: &tenancy.UpdatedAt,
		}
		tenancies = append(tenancies, tenant)
	}
	return tenancies, nil
}

// ServiceName - return service name
func (u UnitServices) ServiceName() string {
	return "UnitServices"
}
