package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	log "github.com/sirupsen/logrus"
)

// ListngService - represent listing services
type ListingService interface {
	GetNearByUnits(ctx context.Context, input *model.NearByUnitsInput) ([]*model.Unit, error)
	ServiceName() string
}

type listingClient struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

func NewListingService(queries *sqlStore.Queries, logger *log.Logger) ListingService {
	return &listingClient{queries: queries, logger: logger}
}

func (l *listingClient) GetNearByUnits(ctx context.Context, input *model.NearByUnitsInput) ([]*model.Unit, error) {
	var units []*model.Unit
	p := fmt.Sprintf("SRID=4326;POINT(%.8f %.8f)", input.Gps.Lng, input.Gps.Lat)

	foundUnits, err := l.queries.GetNearByUnits(ctx, p)
	if err != nil {
		if err == sql.ErrNoRows {
			return units, nil
		} else {
			l.logger.Errorf("%s:%v", l.ServiceName(), err)
			return nil, err
		}
	}

	for _, unit := range foundUnits {
		distance := unit.Distance
		var d string
		dMtr, err := strconv.Atoi(strconv.FormatFloat(distance, 'g', -1, 64))
		if err != nil {
			l.logger.Errorf("%s:%v", l.ServiceName(), err)
			return nil, err
		}

		if dMtr > 1000 {
			d = fmt.Sprintf("%d km", dMtr/1000)
		} else {
			d = fmt.Sprintf("%d m", dMtr)
		}

		u := &model.Unit{
			ID:         unit.ID,
			Name:       unit.Name,
			Price:      strconv.FormatInt(int64(unit.Price), 10),
			Type:       unit.Type,
			Distance:   &d,
			PropertyID: unit.PropertyID.UUID,
			UpdatedAt:  &unit.UpdatedAt,
		}
		units = append(units, u)
	}

	return units, err
}

func (l *listingClient) ServiceName() string {
	return "listingClient"
}
