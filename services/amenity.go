package services

import (
	"context"
	"database/sql"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type AmenityServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// NewAmenityService - factory for amenity services
func NewAmenityService(queries *sqlStore.Queries, logger *log.Logger) *AmenityServices {
	return &AmenityServices{queries, logger}
}

// _ - AmenityServices{} implements AmenityService
var _ interfaces.AmenityService = &AmenityServices{}

// AddAmenity - add property amenity(s)
func (a *AmenityServices) AddAmenity(ctx context.Context, unitID uuid.UUID, amenity *model.UnitAmenityInput) (*model.Amenity, error) {
	insertedAmenity, err := a.queries.CreateAmenity(ctx, sqlStore.CreateAmenityParams{
		Name:     amenity.Name,
		Category: amenity.Category,
		UnitID:   unitID,
	})
	if err != nil {
		a.logger.Errorf("%s:%v", a.ServiceName(), err)
		return nil, err
	}

	return &model.Amenity{
		ID:        insertedAmenity.ID,
		Name:      insertedAmenity.Name,
		Category:  insertedAmenity.Category,
		CreatedAt: &insertedAmenity.CreatedAt,
		UpdatedAt: &insertedAmenity.UpdatedAt,
	}, nil
}

// GetUnitAmenities - grab unit amenities
func (a *AmenityServices) GetUnitAmenities(ctx context.Context, unitID uuid.UUID) ([]*model.Amenity, error) {
	var amenities []*model.Amenity

	foundAmenities, err := a.queries.GetUnitAmenities(ctx, unitID)
	if err != nil {
		if err == sql.ErrNoRows {
			return amenities, nil
		} else {
			a.logger.Errorf("%s:%v", a.ServiceName(), err)
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

// ServiceName - get service name
func (a AmenityServices) ServiceName() string {
	return "AmenityService"
}
