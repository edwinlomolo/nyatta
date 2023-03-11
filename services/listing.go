package services

import (
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	log "github.com/sirupsen/logrus"
)

type ListingServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

var _ interfaces.ListingService = &ListingServices{}

func NewListingService(queries *sqlStore.Queries, logger *log.Logger) *ListingServices {
	return &ListingServices{queries: queries, logger: logger}
}

func (l ListingServices) ServiceName() string {
	return "ListingServices"
}

func (l ListingServices) GetListings(input model.ListingsInput) ([]model.Property, error) {
	foundProperties := make([]model.Property, 0)
	return foundProperties, nil
}
