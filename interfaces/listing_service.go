package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
)

type ListingService interface {
	GetNearByUnits(ctx context.Context, input *model.NearByUnitsInput) ([]*model.Unit, error)
	ServiceName() string
}
