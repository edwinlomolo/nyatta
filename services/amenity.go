package services

import (
	"context"
	"database/sql"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// AmenityService - represent amenity service
type AmenityService interface {
	AddAmenity(ctx context.Context, unitID uuid.UUID, input *model.UnitAmenityInput) (*model.Amenity, error)
	GetUnitAmenities(ctx context.Context, unitID uuid.UUID) ([]*model.Amenity, error)
	ServiceName() string
}

// NewAmenityService - factory for amenity services
func NewAmenityService(queries *sqlStore.Queries, logger *log.Logger) AmenityService {
	return &amenityClient{queries: queries, logger: logger}
}

type amenityClient struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// AddAmenity - add property amenity(s)
func (a *amenityClient) AddAmenity(ctx context.Context, unitID uuid.UUID, amenity *model.UnitAmenityInput) (*model.Amenity, error) {
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
func (a *amenityClient) GetUnitAmenities(ctx context.Context, unitID uuid.UUID) ([]*model.Amenity, error) {
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
func (a *amenityClient) ServiceName() string {
	return "amenityClient"
}
