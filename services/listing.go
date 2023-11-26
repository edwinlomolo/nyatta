package services

import (
	"context"
	"database/sql"
	"fmt"
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

func (l *ListingServices) GetNearByUnits(ctx context.Context, input *model.NearByUnitsInput) ([]*model.Unit, error) {
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
		distance := (unit.Distance).(*string)
		var d *string
		dMtr, err := strconv.Atoi(*distance)
		if err != nil {
			return nil, err
		}

		if dMtr > 1000 {
			*d = fmt.Sprintf("%d km", dMtr/1000)
		} else {
			*d = fmt.Sprintf("%d m", dMtr)
		}

		u := &model.Unit{
			ID:         unit.ID,
			Name:       unit.Name,
			Price:      strconv.FormatInt(int64(unit.Price), 10),
			Distance:   d,
			PropertyID: unit.PropertyID.UUID,
		}
		units = append(units, u)
	}

	return units, err
}

func (l ListingServices) ServiceName() string {
	return "ListingServices"
}
