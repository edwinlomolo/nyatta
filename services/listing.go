package services

import (
	"database/sql"
	"math"
	"strconv"

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
	sqlParams := sqlStore.GetListingsParams{
		Town:     input.Town,
		Type:     "",
		MinPrice: 0,
		MaxPrice: math.MaxInt32,
	}

	if len(*input.PropertyType) != 0 {
		sqlParams.Type = *input.PropertyType
	}
	if *input.MinPrice > 0 {
		sqlParams.MinPrice = int32(*input.MaxPrice)
	}
	if *input.MaxPrice > *input.MinPrice {
		sqlParams.MaxPrice = int32(*input.MaxPrice)
	}

	sqlResult, err := l.queries.GetListings(ctx, sqlParams)
	// Does listing exist?
	if err == sql.ErrNoRows {
		return []model.Property{}, nil
	}

	foundProperties := make([]model.Property, 0)
	for _, foundProperty := range sqlResult {
		property := model.Property{
			ID:         strconv.FormatInt(foundProperty.ID, 10),
			Name:       foundProperty.Name,
			Town:       foundProperty.Town,
			Type:       foundProperty.Type,
			MinPrice:   int(foundProperty.MinPrice),
			MaxPrice:   int(foundProperty.MaxPrice),
			PostalCode: foundProperty.PostalCode,
			CreatedAt:  &foundProperty.CreatedAt,
			UpdatedAt:  &foundProperty.UpdatedAt,
		}
		foundProperties = append(foundProperties, property)
	}
	return foundProperties, nil
}
