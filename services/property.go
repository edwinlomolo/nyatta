package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/cridenour/go-postgis"
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
func (p *PropertyServices) CreateProperty(property *model.NewProperty, createdBy string) (*model.Property, error) {
	gps := postgis.PointS{
		SRID: 4326,
		X:    property.Location.Lat,
		Y:    property.Location.Lng,
	}

	creator, err := strconv.ParseInt(createdBy, 10, 64)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	insertedProperty, err := p.queries.CreateProperty(ctx, sqlStore.CreatePropertyParams{
		Name:      property.Name,
		Type:      property.Type,
		CreatedBy: sql.NullInt64{Int64: creator, Valid: true},
		Location:  fmt.Sprintf("SRID=4326;POINT(%.8f %.8f)", gps.Y, gps.X),
	})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	return &model.Property{
		ID:   strconv.FormatInt(insertedProperty.ID, 10),
		Name: insertedProperty.Name,
		Type: insertedProperty.Type,
		Location: &model.Gps{
			Lat: insertedProperty.Location.X,
			Lng: insertedProperty.Location.Y,
		},
		CreatedAt: &insertedProperty.CreatedAt,
		UpdatedAt: &insertedProperty.UpdatedAt,
	}, nil
}

// GetProperty - return existing property given property id
func (p *PropertyServices) GetProperty(id string) (*model.Property, error) {
	propertyId, err := strconv.Atoi(id)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	foundProperty, err := p.queries.GetProperty(ctx, int64(propertyId))
	if err == sql.ErrNoRows {
		return nil, errors.New("Can't find property")
	}

	return &model.Property{
		ID:        strconv.FormatInt(foundProperty.ID, 10),
		Name:      foundProperty.Name,
		Type:      foundProperty.Type,
		CreatedAt: &foundProperty.CreatedAt,
		UpdatedAt: &foundProperty.UpdatedAt,
	}, nil
}

// PropertiesCreatedBy - get property(s) created by user
func (p *PropertyServices) PropertiesCreatedBy(createdBy string) ([]*model.Property, error) {
	var userProperties []*model.Property

	// Use int64 id
	creator, err := strconv.ParseInt(createdBy, 10, 64)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	props, err := p.queries.PropertiesCreatedBy(ctx, sql.NullInt64{Int64: creator, Valid: true})
	if err == sql.ErrNoRows {
		return userProperties, nil
	}

	for _, item := range props {
		property := &model.Property{
			ID:   strconv.FormatInt(item.ID, 10),
			Name: item.Name,
			Type: item.Type,
			Location: &model.Gps{
				Lat: item.Location.X,
				Lng: item.Location.Y,
			},
			CreatedAt: &item.CreatedAt,
			UpdatedAt: &item.UpdatedAt,
		}
		userProperties = append(userProperties, property)
	}

	return userProperties, nil
}

// GetPropertyUnits - get property units
func (p *PropertyServices) GetPropertyUnits(propertyId string) ([]*model.PropertyUnit, error) {
	var units []*model.PropertyUnit

	id, err := strconv.ParseInt(propertyId, 10, 64)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	foundUnits, err := p.queries.GetPropertyUnits(ctx, sql.NullInt64{Int64: id, Valid: true})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	for _, foundUnit := range foundUnits {
		unit := &model.PropertyUnit{
			ID:         strconv.FormatInt(foundUnit.ID, 10),
			Name:       foundUnit.Name,
			State:      model.UnitState(foundUnit.State),
			Type:       foundUnit.Type,
			PropertyID: strconv.FormatInt(foundUnit.PropertyID.Int64, 10),
			Price:      strconv.FormatInt(int64(foundUnit.Price), 10),
			Bathrooms:  int(foundUnit.Bathrooms),
			CreatedAt:  &foundUnit.CreatedAt,
			UpdatedAt:  &foundUnit.UpdatedAt,
		}
		units = append(units, unit)
	}

	return units, nil
}

// CaretakerVerification - verify caretaker
func (p *PropertyServices) CaretakerPhoneVerification(input *model.CaretakerVerificationInput) (*model.Status, error) {
	status, err := p.twilio.VerifyCode(input.Phone, input.VerifyCode)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}
	return &model.Status{Success: status}, nil
}

// ListingOverview - get listing summary
func (p *PropertyServices) ListingOverview(propertyId string) (*model.ListingOverview, error) {
	id, err := strconv.ParseInt(propertyId, 10, 64)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	totalUnits, err := p.queries.PropertyUnitsCount(ctx, sql.NullInt64{Int64: id, Valid: true})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	occupiedUnits, err := p.queries.OccupiedUnitsCount(ctx, sql.NullInt64{Int64: id, Valid: true})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	vacantUnits, err := p.queries.VacantUnitsCount(ctx, sql.NullInt64{Int64: id, Valid: true})
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
