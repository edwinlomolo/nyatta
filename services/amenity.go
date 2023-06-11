package services

import (
	"errors"
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	log "github.com/sirupsen/logrus"
)

type AmenityServices struct {
	queries         *sqlStore.Queries
	logger          *log.Logger
	propertyService *PropertyServices
}

// NewAmenityService - factory for amenity services
func NewAmenityService(queries *sqlStore.Queries, logger *log.Logger) *AmenityServices {
	propertyService := NewPropertyService(queries, logger)
	return &AmenityServices{queries, logger, propertyService}
}

// _ - AmenityServices{} implements AmenityService
var _ interfaces.AmenityService = &AmenityServices{}

// AddAmenity - add property amenity(s)
func (a *AmenityServices) AddAmenity(amenity *model.AmenityInput) (*model.Amenity, error) {
	// Property exists
	_, err := a.propertyService.GetProperty(amenity.PropertyID)
	if err != nil && err.Error() == "Property does not exist" {
		return nil, errors.New("Adding amenity to non-existent property")
	}

	creator, err := strconv.ParseInt(amenity.PropertyID, 10, 64)
	if err != nil {
		return nil, err
	}
	insertedAmenity, err := a.queries.CreateAmenity(ctx, sqlStore.CreateAmenityParams{
		Name:       amenity.Name,
		Provider:   amenity.Provider,
		Category:   amenity.Category,
		PropertyID: creator,
	})
	if err != nil {
		return nil, err
	}

	return &model.Amenity{
		ID:         strconv.FormatInt(insertedAmenity.ID, 10),
		Name:       insertedAmenity.Name,
		Provider:   insertedAmenity.Provider,
		Category:   insertedAmenity.Category,
		PropertyID: strconv.FormatInt(insertedAmenity.PropertyID, 10),
		CreatedAt:  &insertedAmenity.CreatedAt,
		UpdatedAt:  &insertedAmenity.UpdatedAt,
	}, nil
}

// PropertyAmenities - get property amenities
func (a *AmenityServices) PropertyAmenities(propertyId string) ([]*model.Amenity, error) {
	var amenities []*model.Amenity
	id, err := strconv.ParseInt(propertyId, 10, 64)
	if err != nil {
		return nil, err
	}

	foundAmenities, err := a.queries.PropertyAmenities(ctx, id)
	for _, amenity := range foundAmenities {
		amenities = append(amenities, &model.Amenity{
			ID:         strconv.FormatInt(amenity.ID, 10),
			Name:       amenity.Name,
			Provider:   amenity.Provider,
			Category:   amenity.Category,
			PropertyID: strconv.FormatInt(amenity.PropertyID, 10),
			CreatedAt:  &amenity.CreatedAt,
			UpdatedAt:  &amenity.UpdatedAt,
		})
	}

	return amenities, nil
}
