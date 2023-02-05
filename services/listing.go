package services

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
)

type ListingServices struct {
}

var _ interfaces.ListingService = &ListingServices{}

func NewListingService() *ListingServices {
	return &ListingServices{}
}

func (l ListingServices) ServiceName() string {
	return "Listing"
}

func (l ListingServices) GetListings(model.ListingsInput) ([]model.Property, error) {
	return make([]model.Property, 0), nil
}
