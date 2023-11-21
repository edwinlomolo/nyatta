package services

import (
	"context"
	"database/sql"
	"errors"
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

var (
	ctx context.Context = context.Background()
)

// PropertyServices - represents property service
type PropertyServices struct {
	queries   *sqlStore.Queries
	logger    *log.Logger
	twilio    *TwilioServices
	sendEmail SendEmail
	env       string
}

// _ - PropertyServices{} implements PropertyService
var _ interfaces.PropertyService = &PropertyServices{}

// NewPropertyService - factory for property services
func NewPropertyService(queries *sqlStore.Queries, env string, logger *log.Logger, twilio *TwilioServices, sendEmail SendEmail) *PropertyServices {
	return &PropertyServices{queries: queries, logger: logger, twilio: twilio, sendEmail: sendEmail, env: env}
}

// ServiceName - return service name
func (p PropertyServices) ServiceName() string {
	return "PropertyServices"
}

// CreateProperty - create new property
func (p *PropertyServices) CreateProperty(ctx context.Context, property *model.NewProperty, createdBy uuid.UUID) (*model.Property, error) {
	var caretaker store.Caretaker
	var caretakerErr error
	isLandlord := ctx.Value("is_landlord").(bool)
	phone := ctx.Value("phone").(string)
	gps := postgis.PointS{
		SRID: 4326,
		X:    property.Location.Lat,
		Y:    property.Location.Lng,
	}

	caretaker, caretakerErr = p.queries.GetCaretaker(ctx, property.Caretaker.Phone)
	if caretakerErr != nil && caretakerErr == sql.ErrNoRows {
		caretaker, caretakerErr = p.queries.CreateCaretaker(ctx, sqlStore.CreateCaretakerParams{
			FirstName: property.Caretaker.FirstName,
			LastName:  property.Caretaker.LastName,
			Phone:     property.Caretaker.Phone,
		})
		if caretakerErr != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), caretakerErr)
			return nil, caretakerErr
		}

		if _, err := p.queries.CreateCaretakerAvatar(ctx, sqlStore.CreateCaretakerAvatarParams{
			Upload:      property.Caretaker.Image,
			Category:    model.UploadCategoryProfileImg.String(),
			CaretakerID: uuid.NullUUID{UUID: caretaker.ID, Valid: true},
		}); err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	insertedProperty, err := p.queries.CreateProperty(ctx, sqlStore.CreatePropertyParams{
		Name:        property.Name,
		Type:        property.Type,
		CreatedBy:   uuid.NullUUID{UUID: createdBy, Valid: true},
		Location:    fmt.Sprintf("SRID=4326;POINT(%.8f %.8f)", gps.Y, gps.X),
		CaretakerID: uuid.NullUUID{UUID: caretaker.ID, Valid: true},
	})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	if _, err := p.queries.CreatePropertyThumbnail(ctx, sqlStore.CreatePropertyThumbnailParams{
		Upload:     property.Thumbnail,
		Category:   model.UploadCategoryPropertyThumbnail.String(),
		PropertyID: uuid.NullUUID{UUID: insertedProperty.ID, Valid: true},
	}); err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	if !isLandlord {
		if _, err := p.queries.TrackSubscribeRetries(ctx, sqlStore.TrackSubscribeRetriesParams{
			Phone:            phone,
			SubscribeRetries: 1,
		}); err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	return &model.Property{
		ID:   insertedProperty.ID,
		Name: insertedProperty.Name,
		Type: model.PropertyType(insertedProperty.Type),
		Location: &model.Gps{
			Lat: insertedProperty.Location.X,
			Lng: insertedProperty.Location.Y,
		},
		CreatedBy: insertedProperty.CreatedBy.UUID,
		CreatedAt: &insertedProperty.CreatedAt,
		UpdatedAt: &insertedProperty.UpdatedAt,
	}, nil
}

// GetProperty - return existing property given property id
func (p *PropertyServices) GetProperty(id uuid.UUID) (*model.Property, error) {
	foundProperty, err := p.queries.GetProperty(ctx, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("Can't find property")
	}

	return &model.Property{
		ID:        foundProperty.ID,
		Name:      foundProperty.Name,
		Type:      model.PropertyType(foundProperty.Type),
		CreatedBy: foundProperty.CreatedBy.UUID,
		Location: &model.Gps{
			Lat: foundProperty.Location.X,
			Lng: foundProperty.Location.Y,
		},
		CreatedAt: &foundProperty.CreatedAt,
		UpdatedAt: &foundProperty.UpdatedAt,
	}, nil
}

// PropertiesCreatedBy - get property(s) created by user
func (p *PropertyServices) PropertiesCreatedBy(createdBy uuid.UUID) ([]*model.Property, error) {
	var userProperties []*model.Property

	properties, err := p.queries.PropertiesCreatedBy(ctx, uuid.NullUUID{UUID: createdBy, Valid: true})
	if err == sql.ErrNoRows {
		return userProperties, nil
	}

	for _, item := range properties {
		property := &model.Property{
			ID:   item.ID,
			Name: item.Name,
			Type: model.PropertyType(item.Type),
			Location: &model.Gps{
				Lat: item.Location.X,
				Lng: item.Location.Y,
			},
			CreatedBy: item.CreatedBy.UUID,
			CreatedAt: &item.CreatedAt,
			UpdatedAt: &item.UpdatedAt,
		}
		userProperties = append(userProperties, property)
	}

	return userProperties, nil
}

// GetPropertyUnits - get property units
func (p *PropertyServices) GetPropertyUnits(propertyID uuid.UUID) ([]*model.PropertyUnit, error) {
	var units []*model.PropertyUnit

	foundUnits, err := p.queries.GetPropertyUnits(ctx, uuid.NullUUID{UUID: propertyID, Valid: true})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	for _, foundUnit := range foundUnits {
		unit := &model.PropertyUnit{
			ID:         foundUnit.ID,
			Name:       foundUnit.Name,
			State:      model.UnitState(foundUnit.State),
			Type:       foundUnit.Type,
			PropertyID: foundUnit.PropertyID.UUID,
			Price:      strconv.FormatInt(int64(foundUnit.Price), 10),
			Bathrooms:  int(foundUnit.Bathrooms),
			CreatedAt:  &foundUnit.CreatedAt,
			UpdatedAt:  &foundUnit.UpdatedAt,
		}
		units = append(units, unit)
	}

	return units, nil
}

// CaretakerPhoneVerification - verify caretaker
func (p *PropertyServices) CaretakerPhoneVerification(input *model.CaretakerVerificationInput) (*model.Status, error) {
	status, err := p.twilio.VerifyCode(input.Phone, input.VerifyCode)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}
	return &model.Status{Success: status}, nil
}

// ListingOverview - get listing summary
func (p *PropertyServices) ListingOverview(propertyID uuid.UUID) (*model.ListingOverview, error) {
	pUUID := uuid.NullUUID{UUID: propertyID, Valid: true}

	totalUnits, err := p.queries.PropertyUnitsCount(ctx, pUUID)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	occupiedUnits, err := p.queries.OccupiedUnitsCount(ctx, pUUID)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	vacantUnits, err := p.queries.VacantUnitsCount(ctx, pUUID)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}
	return &model.ListingOverview{
		TotalUnits:    int(totalUnits),
		OccupiedUnits: int(occupiedUnits),
		VacantUnits:   int(vacantUnits),
	}, nil
}

// GetPropertyThumbnail - grab thumbnail
func (p *PropertyServices) GetPropertyThumbnail(propertyID uuid.UUID) (*model.AnyUpload, error) {
	foundThumbnail, err := p.queries.GetPropertyThumbnail(ctx, sqlStore.GetPropertyThumbnailParams{
		PropertyID: uuid.NullUUID{UUID: propertyID, Valid: true},
		Category:   model.UploadCategoryPropertyThumbnail.String(),
	})
	if err != nil {
		return nil, err
	}

	return &model.AnyUpload{
		ID:     foundThumbnail.ID,
		Upload: foundThumbnail.Upload,
	}, nil
}

// GetCaretakerAvatar - grab caretaker avatar
func (p *PropertyServices) GetCaretakerAvatar(caretakerID uuid.UUID) (*model.AnyUpload, error) {
	foundAvatar, err := p.queries.GetCaretakerAvatar(ctx, sqlStore.GetCaretakerAvatarParams{
		CaretakerID: uuid.NullUUID{UUID: caretakerID, Valid: true},
		Category:    model.UploadCategoryProfileImg.String(),
	})
	if err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	return &model.AnyUpload{
		ID:     foundAvatar.ID,
		Upload: foundAvatar.Upload,
	}, nil
}
