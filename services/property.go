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

// PropertyServices - represents property service
type PropertyService interface {
	ServiceName() string
	CreateProperty(ctx context.Context, property *model.NewProperty, createdBy uuid.UUID) (*model.Property, error)
	GetProperty(ctx context.Context, id uuid.UUID) (*model.Property, error)
	GetPropertyThumbnail(ctx context.Context, id uuid.UUID) (*model.AnyUpload, error)
	PropertiesCreatedBy(ctx context.Context, createdBy uuid.UUID) ([]*model.Property, error)
	GetUnits(ctx context.Context, propertyId uuid.UUID) ([]*model.Unit, error)
	CaretakerPhoneVerification(context.Context, *model.CaretakerVerificationInput) (*model.Status, error)
	GetCaretakerAvatar(ctx context.Context, caretakerID uuid.UUID) (*model.AnyUpload, error)
	ListingOverview(ctx context.Context, propertyId uuid.UUID) (*model.ListingOverview, error)
	GetPropertyCaretaker(ctx context.Context, caretakerId uuid.UUID) (*model.Caretaker, error)
	GetUnitsCount(ctx context.Context, id uuid.UUID) (int64, error)
}

type propertyClient struct {
	queries *sqlStore.Queries
	logger  *log.Logger
	twilio  TwilioService
}

// NewPropertyService - factory for property services
func NewPropertyService(queries *sqlStore.Queries, logger *log.Logger, twilio TwilioService) PropertyService {
	return &propertyClient{queries: queries, logger: logger, twilio: twilio}
}

// ServiceName - return service name
func (p *propertyClient) ServiceName() string {
	return "propertyClient"
}

// CreateProperty - create new property
func (p *propertyClient) CreateProperty(ctx context.Context, property *model.NewProperty, createdBy uuid.UUID) (*model.Property, error) {
	var caretaker store.Caretaker
	var caretakerErr error
	phone := ctx.Value("phone").(string)
	userId := ctx.Value("userId").(string)
	gps := postgis.PointS{
		SRID: 4326,
		X:    property.Location.Lat,
		Y:    property.Location.Lng,
	}

	if !property.IsCaretaker {
		caretaker, caretakerErr = p.queries.GetCaretakerByPhone(ctx, property.Caretaker.Phone)
		if caretakerErr != nil && caretakerErr == sql.ErrNoRows {
			caretaker, caretakerErr = p.queries.CreateCaretaker(ctx, sqlStore.CreateCaretakerParams{
				FirstName: property.Caretaker.FirstName,
				LastName:  property.Caretaker.LastName,
				Phone:     property.Caretaker.Phone,
				CreatedBy: uuid.NullUUID{UUID: uuid.MustParse(userId), Valid: true},
			})
			if caretakerErr != nil {
				p.logger.Errorf("%s:%v", p.ServiceName(), caretakerErr)
				return nil, caretakerErr
			}
		}
	} else {
		caretaker, caretakerErr = p.queries.GetCaretakerByPhone(ctx, phone)
		if caretakerErr != nil && caretakerErr == sql.ErrNoRows {
			user, err := ctx.Value("userService").(UserService).GetUser(ctx, uuid.MustParse(userId))
			if err != nil {
				p.logger.Errorf("%s:%v", p.ServiceName(), err)
				return nil, err
			}
			caretaker, caretakerErr = p.queries.CreateCaretaker(ctx, sqlStore.CreateCaretakerParams{
				FirstName: *user.FirstName,
				LastName:  *user.LastName,
				Phone:     phone,
				CreatedBy: uuid.NullUUID{UUID: user.ID, Valid: true},
			})
			if caretakerErr != nil {
				p.logger.Errorf("%s:%v", p.ServiceName(), err)
				return nil, err
			}
		}
	}

	if _, err := p.queries.CreateCaretakerAvatar(ctx, sqlStore.CreateCaretakerAvatarParams{
		Upload:      property.Caretaker.Image,
		Category:    model.UploadCategoryProfileImg.String(),
		CaretakerID: uuid.NullUUID{UUID: caretaker.ID, Valid: true},
	}); err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
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

	return &model.Property{
		ID:          insertedProperty.ID,
		Name:        insertedProperty.Name,
		Type:        model.PropertyType(insertedProperty.Type),
		CaretakerID: &insertedProperty.CaretakerID.UUID,
		CreatedBy:   insertedProperty.CreatedBy.UUID,
		CreatedAt:   &insertedProperty.CreatedAt,
		UpdatedAt:   &insertedProperty.UpdatedAt,
	}, nil
}

// GetProperty - return existing property given property id
func (p *propertyClient) GetProperty(ctx context.Context, id uuid.UUID) (*model.Property, error) {
	foundProperty, err := p.queries.GetProperty(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	var location *model.Point
	json.Unmarshal([]byte((foundProperty.Location).(string)), &location)
	lat := &location.Coordinates[1]
	lng := &location.Coordinates[0]

	return &model.Property{
		ID:          foundProperty.ID,
		Name:        foundProperty.Name,
		CaretakerID: &foundProperty.CaretakerID.UUID,
		Type:        model.PropertyType(foundProperty.Type),
		Location:    &model.Gps{Lat: *lat, Lng: *lng},
		CreatedBy:   foundProperty.CreatedBy.UUID,
		CreatedAt:   &foundProperty.CreatedAt,
		UpdatedAt:   &foundProperty.UpdatedAt,
	}, nil
}

// PropertiesCreatedBy - get property(s) created by user
func (p *propertyClient) PropertiesCreatedBy(ctx context.Context, createdBy uuid.UUID) ([]*model.Property, error) {
	var userProperties []*model.Property

	properties, err := p.queries.PropertiesCreatedBy(ctx, uuid.NullUUID{UUID: createdBy, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return userProperties, nil
		} else {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	for _, item := range properties {
		var location *model.Point
		var gps *model.Gps
		if item.Location != nil {
			json.Unmarshal([]byte((item.Location).(string)), &location)
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
		property := &model.Property{
			ID:          item.ID,
			Name:        item.Name,
			Location:    gps,
			CaretakerID: &item.CaretakerID.UUID,
			Type:        model.PropertyType(item.Type),
			CreatedBy:   item.CreatedBy.UUID,
			CreatedAt:   &item.CreatedAt,
			UpdatedAt:   &item.UpdatedAt,
		}
		userProperties = append(userProperties, property)
	}

	return userProperties, nil
}

// GetUnits - get property units
func (p *propertyClient) GetUnits(ctx context.Context, propertyID uuid.UUID) ([]*model.Unit, error) {
	var units []*model.Unit

	foundUnits, err := p.queries.GetPropertyUnits(ctx, uuid.NullUUID{UUID: propertyID, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return units, nil
		} else {
			p.logger.Errorf("%s: %v", p.ServiceName(), err)
			return nil, err
		}
	}

	for _, foundUnit := range foundUnits {
		unit := &model.Unit{
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
func (p *propertyClient) CaretakerPhoneVerification(ctx context.Context, input *model.CaretakerVerificationInput) (*model.Status, error) {
	status, err := p.twilio.VerifyCode(input.Phone, input.VerifyCode)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}
	return &model.Status{Success: status}, nil
}

// ListingOverview - get listing summary
func (p *propertyClient) ListingOverview(ctx context.Context, propertyID uuid.UUID) (*model.ListingOverview, error) {
	pUUID := uuid.NullUUID{UUID: propertyID, Valid: true}

	totalUnits, err := p.queries.UnitsCount(ctx, pUUID)
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
func (p *propertyClient) GetPropertyThumbnail(ctx context.Context, propertyID uuid.UUID) (*model.AnyUpload, error) {
	foundThumbnail, err := p.queries.GetPropertyThumbnail(ctx, sqlStore.GetPropertyThumbnailParams{
		PropertyID: uuid.NullUUID{UUID: propertyID, Valid: true},
		Category:   model.UploadCategoryPropertyThumbnail.String(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	return &model.AnyUpload{
		ID:     foundThumbnail.ID,
		Upload: foundThumbnail.Upload,
	}, nil
}

// GetCaretakerAvatar - grab caretaker avatar
func (p *propertyClient) GetCaretakerAvatar(ctx context.Context, caretakerID uuid.UUID) (*model.AnyUpload, error) {
	foundAvatar, err := p.queries.GetCaretakerAvatar(ctx, sqlStore.GetCaretakerAvatarParams{
		CaretakerID: uuid.NullUUID{UUID: caretakerID, Valid: true},
		Category:    model.UploadCategoryProfileImg.String(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	return &model.AnyUpload{
		ID:     foundAvatar.ID,
		Upload: foundAvatar.Upload,
	}, nil
}

// GetPropertyCaretaker - grab property caretaker
func (p *propertyClient) GetPropertyCaretaker(ctx context.Context, caretakerID uuid.UUID) (*model.Caretaker, error) {
	foundCaretaker, err := p.queries.GetCaretakerById(ctx, caretakerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	return &model.Caretaker{
		ID:        foundCaretaker.ID,
		FirstName: foundCaretaker.FirstName,
		LastName:  foundCaretaker.LastName,
		Phone:     foundCaretaker.Phone,
		CreatedAt: &foundCaretaker.CreatedAt,
		UpdatedAt: &foundCaretaker.UpdatedAt,
	}, nil
}

// GetUnitsCount - grab total property units
func (p *propertyClient) GetUnitsCount(ctx context.Context, id uuid.UUID) (int64, error) {
	totalCount, err := p.queries.UnitsCount(ctx, uuid.NullUUID{UUID: id, Valid: true})
	if err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return 0, err
	}

	return totalCount, nil
}
