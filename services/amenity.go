package services

import (
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
func NewAmenityService(queries *sqlStore.Queries, logger *log.Logger, propertyService *PropertyServices) *AmenityServices {
	return &AmenityServices{queries, logger, propertyService}
}

// _ - AmenityServices{} implements AmenityService
var _ interfaces.AmenityService = &AmenityServices{}

// AddAmenity - add property amenity(s)
func (a *AmenityServices) AddAmenity(amenity *model.UnitAmenityInput) (*model.Amenity, error) {
	return &model.Amenity{}, nil
}
