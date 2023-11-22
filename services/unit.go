package services

import (
	"database/sql"
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

	if len(input.Bedrooms) > 0 {
		for _, bedroom := range input.Bedrooms {
			if _, err := u.queries.CreateUnitBedroom(ctx, sqlStore.CreateUnitBedroomParams{
				PropertyUnitID: unit.ID,
				BedroomNumber:  int32(bedroom.BedroomNumber),
				EnSuite:        bedroom.EnSuite,
				Master:         bedroom.Master,
			}); err != nil {
				u.logger.Errorf("%s: %v", u.ServiceName(), err)
				return nil, err
			}
		}
	}

	if len(input.Amenities) > 0 {
		for _, amenity := range input.Amenities {
			if _, err := u.queries.CreateAmenity(ctx, sqlStore.CreateAmenityParams{
				Name:           amenity.Name,
				Category:       amenity.Category,
				PropertyUnitID: unit.ID,
			}); err != nil {
				u.logger.Errorf("%s: %v", u.ServiceName(), err)
				return nil, err
			}
		}
	}

	uidUUID := uuid.NullUUID{UUID: unit.ID, Valid: true}
	for _, upload := range input.Uploads {
		if _, err := u.queries.CreateUnitImage(ctx, sqlStore.CreateUnitImageParams{
			PropertyUnitID: uidUUID,
			Category:       model.UploadCategoryUnitImages.String(),
			Label:          sql.NullString{String: upload.Category, Valid: true},
			Upload:         upload.Image,
		}); err != nil {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	return &model.PropertyUnit{
		ID:         unit.ID,
		Name:       unit.Name,
		Bathrooms:  int(unit.Bathrooms),
		PropertyID: input.PropertyID,
		CreatedAt:  &unit.CreatedAt,
		UpdatedAt:  &unit.UpdatedAt,
		Price:      strconv.FormatInt(int64(unit.Price), 10),
		Type:       unit.Type,
	}, nil
}

// GetUnitBedrooms - return unit bedrooms
func (u *UnitServices) GetUnitBedrooms(unitId uuid.UUID) ([]*model.Bedroom, error) {
	var bedrooms []*model.Bedroom

	foundBedrooms, err := u.queries.GetUnitBedrooms(ctx, unitId)
	if err != nil {
		if err == sql.ErrNoRows {
			return bedrooms, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	for _, unit := range foundBedrooms {
		bedroom := &model.Bedroom{
			ID:            unitId,
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

// ServiceName - return service name
func (u UnitServices) ServiceName() string {
	return "UnitServices"
}

// GetUnitImages - grab images
func (u *UnitServices) GetUnitImages(id uuid.UUID) ([]*model.AnyUpload, error) {
	var images []*model.AnyUpload

	foundUploads, err := u.queries.GetUnitImages(ctx, sqlStore.GetUnitImagesParams{
		PropertyUnitID: uuid.NullUUID{UUID: id, Valid: true},
		Category:       model.UploadCategoryUnitImages.String(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return images, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	for _, upload := range foundUploads {
		image := &model.AnyUpload{
			ID:       upload.ID,
			Category: upload.Label.String,
			Upload:   upload.Upload,
		}
		images = append(images, image)
	}

	return images, nil
}

// GetUnitAmenities - grab unit amenities
func (u *UnitServices) GetUnitAmenities(unitID uuid.UUID) ([]*model.Amenity, error) {
	var amenities []*model.Amenity

	foundAmenities, err := u.queries.GetUnitAmenities(ctx, unitID)
	if err != nil {
		if err == sql.ErrNoRows {
			return amenities, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	for _, amenity := range foundAmenities {
		unitAmenity := &model.Amenity{
			ID:       amenity.ID,
			Category: amenity.Category,
			Name:     amenity.Name,
		}
		amenities = append(amenities, unitAmenity)
	}

	return amenities, nil
}

// GetPropertyUnit - grab unit
func (u *UnitServices) GetPropertyUnit(unitID uuid.UUID) (*model.PropertyUnit, error) {
	foundUnit, err := u.queries.GetPropertyUnit(ctx, unitID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
		}
	}

	return &model.PropertyUnit{
		ID:         foundUnit.ID,
		Name:       foundUnit.Name,
		State:      model.UnitState(foundUnit.State),
		PropertyID: foundUnit.PropertyID.UUID,
		Type:       foundUnit.Type,
		Price:      strconv.FormatInt(int64(foundUnit.Price), 10),
		Bathrooms:  int(foundUnit.Bathrooms),
		CreatedAt:  &foundUnit.CreatedAt,
		UpdatedAt:  &foundUnit.UpdatedAt,
	}, nil
}
