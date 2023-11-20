package services

import (
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/google/uuid"
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
	pUUID := uuid.NullUUID{UUID: input.PropertyID, Valid: true}

	unitPrice, err := strconv.ParseInt(input.Price, 10, 64)
	if err != nil {
		u.logger.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}
	unit, err := u.queries.CreatePropertyUnit(ctx, sqlStore.CreatePropertyUnitParams{
		PropertyID: pUUID,
		Name:       input.Name,
		Price:      int32(unitPrice),
		Type:       input.Type,
		Bathrooms:  int32(input.Baths),
	})

	// TODO for every uploads
	if _, err := u.queries.CreateUnitImage(ctx, sqlStore.CreateUnitImageParams{}); err != nil {
		u.logger.Errorf("%s:%v", u.ServiceName(), err)
		return nil, err
	}

	if len(input.Bedrooms) > 0 {
		for i := 0; i < len(input.Bedrooms); i++ {
			_, err := u.queries.CreateUnitBedroom(ctx, sqlStore.CreateUnitBedroomParams{
				PropertyUnitID: unit.ID,
				BedroomNumber:  int32(input.Bedrooms[i].BedroomNumber),
				EnSuite:        input.Bedrooms[i].EnSuite,
				Master:         input.Bedrooms[i].Master,
			})
			if err != nil {
				u.logger.Errorf("%s: %v", u.ServiceName(), err)
				return nil, err
			}
		}
	}
	// TODO create unit amenity
	uidUUID := uuid.NullUUID{UUID: unit.ID, Valid: true}
	if len(input.Amenities) > 0 {
		for j := 0; j < len(input.Amenities); j++ {
			_, err := u.queries.CreateAmenity(ctx, sqlStore.CreateAmenityParams{
				Name:           input.Amenities[j].Name,
				Category:       input.Amenities[j].Category,
				PropertyUnitID: uidUUID,
			})
			if err != nil {
				u.logger.Errorf("%s: %v", u.ServiceName(), err)
				return nil, err
			}
		}
	}

	return &model.PropertyUnit{
		ID:         unit.ID,
		Bathrooms:  int(unit.Bathrooms),
		PropertyID: input.PropertyID,
		CreatedAt:  &unit.CreatedAt,
		UpdatedAt:  &unit.UpdatedAt,
		Price:      strconv.FormatInt(int64(unit.Price), 10),
		Type:       unit.Type,
	}, nil
}

// AddUnitBedroom - add property unit bedroom(s)
func (u *UnitServices) AddUnitBedrooms(input []*model.UnitBedroomInput) ([]*model.Bedroom, error) {
	var insertedBedrooms []*model.Bedroom
	return insertedBedrooms, nil
	//for _, value := range input {
	//	if value.BedroomNumber == 0 {
	//		return nil, errors.New("Zero is not a valid value")
	//	}
	//	propertyId := *value.PropertyUnitID
	//	propertyUnitId, err := strconv.ParseInt(propertyId, 10, 64)
	//	if err != nil {
	//		return nil, err
	//	}
	//	insertedBedroom, err := u.queries.CreateUnitBedroom(ctx, sqlStore.CreateUnitBedroomParams{
	//		PropertyUnitID: propertyUnitId,
	//		BedroomNumber:  int32(value.BedroomNumber),
	//		EnSuite:        value.EnSuite,
	//		Master:         value.Master,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	insertedBedrooms = append(insertedBedrooms, &model.Bedroom{
	//		ID:        strconv.FormatInt(insertedBedroom.ID, 10),
	//		EnSuite:   insertedBedroom.EnSuite,
	//		Master:    insertedBedroom.Master,
	//		CreatedAt: &insertedBedroom.CreatedAt,
	//		UpdatedAt: &insertedBedroom.UpdatedAt,
	//	})
	//}
	//return insertedBedrooms, nil
}

// GetUnitBedrooms - return unit bedrooms
func (u *UnitServices) GetUnitBedrooms(unitId uuid.UUID) ([]*model.Bedroom, error) {
	var bedrooms []*model.Bedroom
	return bedrooms, nil
	//id, err := strconv.ParseInt(unitId, 10, 64)
	//if err != nil {
	//	return nil, err
	//}

	//foundBedrooms, err := u.queries.GetUnitBedrooms(ctx, id)
	//for _, unit := range foundBedrooms {
	//	bedroom := &model.Bedroom{
	//		ID:            strconv.FormatInt(unit.ID, 10),
	//		BedroomNumber: int(unit.BedroomNumber),
	//		EnSuite:       unit.EnSuite,
	//		Master:        unit.Master,
	//		CreatedAt:     &unit.CreatedAt,
	//		UpdatedAt:     &unit.UpdatedAt,
	//	}
	//	bedrooms = append(bedrooms, bedroom)
	//}
	//return bedrooms, nil
}

// GetUnitTenancy - return unit tenancy
func (u *UnitServices) GetUnitTenancy(unitId uuid.UUID) ([]*model.Tenant, error) {
	var tenancies []*model.Tenant

	return tenancies, nil
	//id, err := strconv.ParseInt(unitId, 10, 64)
	//if err != nil {
	//	return nil, err
	//}

	//foundTenancies, err := u.queries.GetUnitTenancy(ctx, id)
	//for _, tenancy := range foundTenancies {
	//	tenant := &model.Tenant{
	//		ID:        strconv.FormatInt(tenancy.ID, 10),
	//		StartDate: tenancy.StartDate,
	//		EndDate:   &tenancy.EndDate.Time,
	//		CreatedAt: &tenancy.CreatedAt,
	//		UpdatedAt: &tenancy.UpdatedAt,
	//	}
	//	tenancies = append(tenancies, tenant)
	//}
	//return tenancies, nil
}

// ServiceName - return service name
func (u UnitServices) ServiceName() string {
	return "UnitServices"
}

// GetUnitImages - grab uploads
func (u *UnitServices) GetUnitImages(id uuid.UUID) ([]*model.AnyUpload, error) {
	var images []*model.AnyUpload
	return images, nil
}
