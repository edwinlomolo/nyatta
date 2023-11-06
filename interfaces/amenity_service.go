package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
)

type AmenityService interface {
	AddAmenity(*model.UnitAmenityInput) (*model.Amenity, error)
}
