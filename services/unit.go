package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/3dw1nM0535/nyatta/database/store"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/cridenour/go-postgis"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// UnitService - represent unit services
type UnitService interface {
	AddUnit(ctx context.Context, unit *model.UnitInput) (*model.Unit, error)
	GetUnitBedrooms(ctx context.Context, unitId uuid.UUID) ([]*model.Bedroom, error)
	GetUnitImages(ctx context.Context, id uuid.UUID) ([]*model.AnyUpload, error)
	GetUnit(ctx context.Context, unitID uuid.UUID) (*model.Unit, error)
	ServiceName() string
	UnitsCreatedBy(ctx context.Context, createdBy uuid.UUID) ([]*model.Unit, error)
	GetUnitThumbnail(ctx context.Context, id uuid.UUID) (*model.AnyUpload, error)
}
type unitClient struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// NewUnitService - factory for unitClient
func NewUnitService(queries *sqlStore.Queries, logger *log.Logger) UnitService {
	return &unitClient{queries, logger}
}

// AddUnit - add property unit
func (u *unitClient) AddUnit(ctx context.Context, input *model.UnitInput) (*model.Unit, error) {
	phone := ctx.Value("phone").(string)
	userId := ctx.Value("userId").(string)
	amenityService := ctx.Value("amenityService").(AmenityService)
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
					CreatedBy: uuid.NullUUID{UUID: uuid.MustParse(userId), Valid: true},
				})
				if caretakerErr != nil {
					u.logger.Errorf("%s:%v", u.ServiceName(), caretakerErr)
					return nil, caretakerErr
				}
			}
		} else {
			caretaker, caretakerErr = u.queries.GetCaretakerByPhone(ctx, phone)
			if caretakerErr != nil && caretakerErr == sql.ErrNoRows {
				user, err := ctx.Value("userService").(UserService).GetUser(ctx, uuid.MustParse(userId))
				if err != nil {
					u.logger.Errorf("%s:%v", u.ServiceName(), err)
					return nil, err
				}
				caretaker, caretakerErr = u.queries.CreateCaretaker(ctx, sqlStore.CreateCaretakerParams{
					FirstName: *user.FirstName,
					LastName:  *user.LastName,
					Phone:     phone,
					CreatedBy: uuid.NullUUID{UUID: user.ID, Valid: true},
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

		if _, err := u.queries.CreateCaretakerAvatar(ctx, sqlStore.CreateCaretakerAvatarParams{
			Upload:      input.Caretaker.Image,
			Category:    model.UploadCategoryProfileImg.String(),
			CaretakerID: uuid.NullUUID{UUID: caretaker.ID, Valid: true},
		}); err != nil {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
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
			if _, err := amenityService.AddAmenity(ctx, unit.ID, &model.UnitAmenityInput{
				Name:     amenity.Name,
				Category: amenity.Category,
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
func (u *unitClient) GetUnitBedrooms(ctx context.Context, unitId uuid.UUID) ([]*model.Bedroom, error) {
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
func (u *unitClient) ServiceName() string {
	return "unitClient"
}

// GetUnitThumbnail - just grab one from images upload
func (u *unitClient) GetUnitThumbnail(ctx context.Context, id uuid.UUID) (*model.AnyUpload, error) {
	foundUpload, err := u.queries.GetUnitThumbnail(ctx, sqlStore.GetUnitThumbnailParams{
		UnitID:   uuid.NullUUID{UUID: id, Valid: true},
		Category: model.UploadCategoryUnitImages.String(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	return &model.AnyUpload{
		ID:     foundUpload.ID,
		Upload: foundUpload.Upload,
	}, nil
}

// GetUnitImages - grab images
func (u *unitClient) GetUnitImages(ctx context.Context, id uuid.UUID) ([]*model.AnyUpload, error) {
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

// GetUnit - grab unit
func (u *unitClient) GetUnit(ctx context.Context, unitID uuid.UUID) (*model.Unit, error) {
	foundUnit, err := u.queries.GetUnit(ctx, unitID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	var location *model.Point
	var gps *model.Gps
	if foundUnit.Location != nil {
		json.Unmarshal([]byte((foundUnit.Location).(string)), &location)
	} else {
		location = nil
		gps = nil
	}

	if location != nil {
		lat := &location.Coordinates[1]
		lng := &location.Coordinates[0]
		gps = &model.Gps{
			Lng: *lng,
			Lat: *lat,
		}
	}

	return &model.Unit{
		ID:          foundUnit.ID,
		Name:        foundUnit.Name,
		PropertyID:  foundUnit.PropertyID.UUID,
		Type:        foundUnit.Type,
		CaretakerID: &foundUnit.CaretakerID.UUID,
		Location:    gps,
		Price:       strconv.FormatInt(int64(foundUnit.Price), 10),
		Bathrooms:   int(foundUnit.Bathrooms),
		CreatedAt:   &foundUnit.CreatedAt,
		UpdatedAt:   &foundUnit.UpdatedAt,
	}, nil
}

// UnitsCreatedBy - grab user listings
func (u *unitClient) UnitsCreatedBy(ctx context.Context, createdBy uuid.UUID) ([]*model.Unit, error) {
	var units []*model.Unit
	foundUnits, err := u.queries.UnitsCreatedBy(ctx, uuid.NullUUID{UUID: createdBy, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return units, nil
		} else {
			u.logger.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
	}

	for _, item := range foundUnits {
		unit := &model.Unit{
			ID:          item.ID,
			Name:        item.Name,
			Type:        item.Type,
			CaretakerID: &item.CaretakerID.UUID,
			CreatedBy:   &item.CreatedBy.UUID,
			CreatedAt:   &item.CreatedAt,
			UpdatedAt:   &item.UpdatedAt,
		}
		units = append(units, unit)
	}

	return units, nil
}
