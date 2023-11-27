package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/3dw1nM0535/nyatta/database/store"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/cridenour/go-postgis"
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

// AddUnit - add property unit
func (u *UnitServices) AddUnit(ctx context.Context, input *model.UnitInput) (*model.Unit, error) {
	phone := ctx.Value("phone").(string)
	userId := ctx.Value("userId").(string)
	notUnitType := input.Type != "Unit"
	var caretaker store.Caretaker
	var caretakerErr error

	unitPrice, err := strconv.ParseInt(input.Price, 10, 64)
	if err != nil {
		u.logger.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}

	if notUnitType {
		if *input.IsCaretaker == false {
			caretaker, caretakerErr = u.queries.GetCaretakerByPhone(ctx, input.Caretaker.Phone)
			if caretakerErr != nil && caretakerErr == sql.ErrNoRows {
				caretaker, caretakerErr = u.queries.CreateCaretaker(ctx, sqlStore.CreateCaretakerParams{
					FirstName: input.Caretaker.FirstName,
					LastName:  input.Caretaker.LastName,
					Phone:     input.Caretaker.Phone,
				})
				if caretakerErr != nil {
					u.logger.Errorf("%s:%v", u.ServiceName(), caretakerErr)
					return nil, caretakerErr
				}
			}
		} else {
			caretaker, caretakerErr = u.queries.GetCaretakerByPhone(ctx, phone)
			if caretakerErr != nil && caretakerErr == sql.ErrNoRows {
				user, err := ctx.Value("userService").(*UserServices).GetUser(ctx, uuid.MustParse(userId))
				if err != nil {
					u.logger.Errorf("%s:%v", u.ServiceName(), err)
					return nil, err
				}
				caretaker, caretakerErr = u.queries.CreateCaretaker(ctx, sqlStore.CreateCaretakerParams{
					FirstName: *user.FirstName,
					LastName:  *user.LastName,
					Phone:     phone,
				})
				if caretakerErr != nil {
					u.logger.Errorf("%s:%v", u.ServiceName(), err)
					return nil, err
				}
			}
		}
	}

	var unit store.Unit
	var unitErr error
	if notUnitType {
		gps := postgis.PointS{
			X: input.Location.Lat,
			Y: input.Location.Lng,
		}
		if unit, unitErr = u.queries.CreateOtherUnit(ctx, sqlStore.CreateOtherUnitParams{
			Name:        input.Name,
			State:       input.State.String(),
			Location:    fmt.Sprintf("SRID=4326;POINT(%.8f %.8f)", gps.Y, gps.X),
			Price:       int32(unitPrice),
			Type:        input.Type,
			Bathrooms:   int32(input.Baths),
			CaretakerID: uuid.NullUUID{UUID: caretaker.ID, Valid: true},
			CreatedBy:   uuid.NullUUID{UUID: uuid.MustParse(userId), Valid: true},
		}); unitErr != nil {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, unitErr
		}
	} else {
		pUUID := uuid.NullUUID{UUID: *input.PropertyID, Valid: true}
		if unit, unitErr = u.queries.CreateUnit(ctx, sqlStore.CreateUnitParams{
			PropertyID: pUUID,
			Name:       input.Name,
			State:      input.State.String(),
			Price:      int32(unitPrice),
			Type:       input.Type,
			Bathrooms:  int32(input.Baths),
		}); unitErr != nil {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, unitErr
		}
	}

	if len(input.Bedrooms) > 0 {
		for _, bedroom := range input.Bedrooms {
			if _, err := u.queries.CreateUnitBedroom(ctx, sqlStore.CreateUnitBedroomParams{
				UnitID:        unit.ID,
				BedroomNumber: int32(bedroom.BedroomNumber),
				EnSuite:       bedroom.EnSuite,
				Master:        bedroom.Master,
			}); err != nil {
				u.logger.Errorf("%s: %v", u.ServiceName(), err)
				return nil, err
			}
		}
	}

	if len(input.Amenities) > 0 {
		for _, amenity := range input.Amenities {
			if _, err := u.queries.CreateAmenity(ctx, sqlStore.CreateAmenityParams{
				Name:     amenity.Name,
				Category: amenity.Category,
				UnitID:   unit.ID,
			}); err != nil {
				u.logger.Errorf("%s: %v", u.ServiceName(), err)
				return nil, err
			}
		}
	}

	uidUUID := uuid.NullUUID{UUID: unit.ID, Valid: true}
	for _, upload := range input.Uploads {
		if _, err := u.queries.CreateUnitImage(ctx, sqlStore.CreateUnitImageParams{
			UnitID:   uidUUID,
			Category: model.UploadCategoryUnitImages.String(),
			Label:    sql.NullString{String: upload.Category, Valid: true},
			Upload:   upload.Image,
		}); err != nil {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	return &model.Unit{
		ID:         unit.ID,
		Name:       unit.Name,
		Bathrooms:  int(unit.Bathrooms),
		PropertyID: unit.PropertyID.UUID,
		CreatedAt:  &unit.CreatedAt,
		UpdatedAt:  &unit.UpdatedAt,
		Price:      strconv.FormatInt(int64(unit.Price), 10),
		Type:       unit.Type,
	}, nil
}

// GetUnitBedrooms - return unit bedrooms
func (u *UnitServices) GetUnitBedrooms(ctx context.Context, unitId uuid.UUID) ([]*model.Bedroom, error) {
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
func (u *UnitServices) GetUnitImages(ctx context.Context, id uuid.UUID) ([]*model.AnyUpload, error) {
	var images []*model.AnyUpload

	foundUploads, err := u.queries.GetUnitImages(ctx, sqlStore.GetUnitImagesParams{
		UnitID:   uuid.NullUUID{UUID: id, Valid: true},
		Category: model.UploadCategoryUnitImages.String(),
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
func (u *UnitServices) GetUnitAmenities(ctx context.Context, unitID uuid.UUID) ([]*model.Amenity, error) {
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

// GetUnit - grab unit
func (u *UnitServices) GetUnit(ctx context.Context, unitID uuid.UUID) (*model.Unit, error) {
	foundUnit, err := u.queries.GetUnit(ctx, unitID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	var location *model.Gps
	if foundUnit.Location != nil {
		location = (foundUnit.Location).(*model.Gps)
	} else {
		location = nil
	}

	return &model.Unit{
		ID:          foundUnit.ID,
		Name:        foundUnit.Name,
		PropertyID:  foundUnit.PropertyID.UUID,
		Type:        foundUnit.Type,
		Location:    location,
		CaretakerID: &foundUnit.CaretakerID.UUID,
		Price:       strconv.FormatInt(int64(foundUnit.Price), 10),
		Bathrooms:   int(foundUnit.Bathrooms),
		CreatedAt:   &foundUnit.CreatedAt,
		UpdatedAt:   &foundUnit.UpdatedAt,
	}, nil
}
