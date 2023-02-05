package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
)

type ListingService interface {
	ServiceName() string
	GetListings(model.ListingsInput) ([]model.Property, error)
}
