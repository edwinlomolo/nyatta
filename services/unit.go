package services

import (
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/interfaces"
	log "github.com/sirupsen/logrus"
)

type UnitServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// _ - UnitServices{} implements UnitService interface
var _ interfaces.UnitService = &UnitServices{}

// NewUnitService - factory for UnitServices
func NewUnitService(queries *sqlStore.Queries, logger *log.Logger) *UnitServices {
	return &UnitServices{queries, logger}
}

// ServiceName - return service name
func (u UnitServices) ServiceName() string {
	return "UnitServices"
}
