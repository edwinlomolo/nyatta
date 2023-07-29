package services

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
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
func (p *PropertyServices) CreateProperty(property *model.NewProperty) (*model.Property, error) {
	creator, err := strconv.ParseInt(property.CreatedBy, 10, 64)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}
	insertedProperty, err := p.queries.CreateProperty(ctx, sqlStore.CreatePropertyParams{
		Name:       property.Name,
		Town:       property.Town,
		Type:       property.Type,
		PostalCode: property.PostalCode,
		CreatedBy:  creator,
	})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	return &model.Property{
		ID:         strconv.FormatInt(insertedProperty.ID, 10),
		Name:       insertedProperty.Name,
		Town:       insertedProperty.Town,
		PostalCode: insertedProperty.PostalCode,
		Type:       insertedProperty.Type,
		Status:     insertedProperty.Status,
		CreatedBy:  strconv.FormatInt(creator, 10),
		CreatedAt:  &insertedProperty.CreatedAt,
		UpdatedAt:  &insertedProperty.UpdatedAt,
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
		// TODO just return empty property list
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, errors.New("Property does not exist")
	}
	return &model.Property{
		ID:         strconv.FormatInt(foundProperty.ID, 10),
		Name:       foundProperty.Name,
		Town:       foundProperty.Town,
		Type:       foundProperty.Type,
		Status:     foundProperty.Status,
		PostalCode: foundProperty.PostalCode,
		CreatedBy:  strconv.FormatInt(foundProperty.CreatedBy, 10),
		CreatedAt:  &foundProperty.CreatedAt,
		UpdatedAt:  &foundProperty.UpdatedAt,
	}, nil
}

// FindByTown - find property(s) in a given town
func (p *PropertyServices) FindByTown(town string) ([]*model.Property, error) {
	return make([]*model.Property, 0), nil
}

// FindByPostalCode - find property(s) in a given postal
func (p *PropertyServices) FindByPostalCode(postalCode string) ([]*model.Property, error) {
	return make([]*model.Property, 0), nil
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

	props, err := p.queries.PropertiesCreatedBy(ctx, creator)
	if err == sql.ErrNoRows {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, errors.New("No properties found")
	}

	for _, item := range props {
		property := &model.Property{
			ID:         strconv.FormatInt(item.ID, 10),
			Name:       item.Name,
			Town:       item.Town,
			Type:       item.Type,
			Status:     item.Status,
			PostalCode: item.PostalCode,
			CreatedAt:  &item.CreatedAt,
			UpdatedAt:  &item.UpdatedAt,
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

	foundUnits, err := p.queries.GetPropertyUnits(ctx, id)
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
			PropertyID: strconv.FormatInt(foundUnit.PropertyID, 10),
			Price:      strconv.FormatInt(int64(foundUnit.Price), 10),
			Bathrooms:  int(foundUnit.Bathrooms),
			CreatedAt:  &foundUnit.CreatedAt,
			UpdatedAt:  &foundUnit.UpdatedAt,
		}
		units = append(units, unit)
	}

	return units, nil
}

// SetupProperty - setup listing
func (p PropertyServices) SetupProperty(input *model.SetupPropertyInput) (*model.Status, error) {
	user, err := p.queries.FindByEmail(ctx, sql.NullString{String: input.Creator, Valid: true})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	// Create property caretaker
	caretaker, err := p.queries.CreateCaretaker(ctx, sqlStore.CreateCaretakerParams{
		FirstName:      input.Caretaker.FirstName,
		LastName:       input.Caretaker.LastName,
		Phone:          sql.NullString{String: input.Caretaker.Phone, Valid: true},
		CountryCode:    input.Caretaker.CountryCode.String(),
		Idverification: input.Caretaker.IDVerification,
		Image:          input.Shoot.ContactPerson,
	})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	// Create property
	property, err := p.queries.CreateProperty(ctx, sqlStore.CreatePropertyParams{
		Name:       input.Name,
		Town:       input.Town,
		PostalCode: input.PostalCode,
		Type:       input.PropertyType,
		CreatedBy:  user.ID,
		Caretaker:  sql.NullInt64{Int64: caretaker.ID, Valid: true},
	})
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	// Create property unit
	if len(input.Units) > 0 {
		for i := 0; i < len(input.Units); i++ {
			unitPrice, err := strconv.ParseInt(input.Units[i].Price, 10, 64)
			unit, err := p.queries.CreatePropertyUnit(ctx, sqlStore.CreatePropertyUnitParams{
				Name:       input.Units[i].Name,
				Type:       input.Units[i].Type,
				Bathrooms:  int32(input.Units[i].Baths),
				Price:      int32(unitPrice),
				PropertyID: property.ID,
			})
			// create unit bedrooms
			if len(input.Units[i].Bedrooms) > 0 {
				for j := 0; j < len(input.Units[i].Bedrooms); j++ {
					_, err := p.queries.CreateUnitBedroom(ctx, sqlStore.CreateUnitBedroomParams{
						PropertyUnitID: unit.ID,
						BedroomNumber:  int32(input.Units[i].Bedrooms[j].BedroomNumber),
						EnSuite:        input.Units[i].Bedrooms[j].EnSuite,
						Master:         input.Units[i].Bedrooms[j].Master,
					})
					if err != nil {
						p.logger.Errorf("%s: %v", p.ServiceName(), err)
						return nil, err
					}
				}
			}
			// create unit amenities
			if len(input.Units[i].Amenities) > 0 {
				for j := 0; j < len(input.Units[i].Amenities); j++ {
					_, err := p.queries.CreateAmenity(ctx, sqlStore.CreateAmenityParams{
						Name:           input.Units[i].Amenities[j].Name,
						Category:       input.Units[i].Amenities[j].Category,
						PropertyUnitID: unit.ID,
					})
					if err != nil {
						p.logger.Errorf("%s: %v", p.ServiceName(), err)
						return nil, err
					}
				}
			}
			if err != nil {
				p.logger.Errorf("%s: %v", p.ServiceName(), err)
				return nil, err
			}
		}
	}
	// send email
	if p.env == "staging" || p.env == "production" {
		from := os.Getenv("EMAIL_FROM")
		p.sendEmail([]string{user.Email.String}, from, "Congratulations! Welcome Onboard", newPropertyEmail)
	}
	return &model.Status{Success: "okay"}, nil
}

// CaretakerVerification - verify caretaker
func (p *PropertyServices) CaretakerPhoneVerification(input *model.CaretakerVerificationInput) (*model.Status, error) {
	status, err := p.twilio.VerifyCode(input.Phone, input.VerifyCode, input.CountryCode)
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

	totalUnits, err := p.queries.PropertyUnitsCount(ctx, id)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	occupiedUnits, err := p.queries.OccupiedUnitsCount(ctx, id)
	if err != nil {
		p.logger.Errorf("%s: %v", p.ServiceName(), err)
		return nil, err
	}

	vacantUnits, err := p.queries.VacantUnitsCount(ctx, id)
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
