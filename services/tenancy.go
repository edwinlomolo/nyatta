package services

import (
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/interfaces"
	log "github.com/sirupsen/logrus"
)

type TenancyServices struct {
	queries *sqlStore.Queries
	logger  *log.Logger
}

// _ - TenancyServices{} implements TenancyService interface
var _ interfaces.TenancyService = &TenancyServices{}

// NewTenancyService - factory for tenancy services
func NewTenancyService(queries *sqlStore.Queries, logger *log.Logger) *TenancyServices {
	return &TenancyServices{queries, logger}
}

// ServiceName - return service name
func (t TenancyServices) ServiceName() string {
	return "TenancyServices"
}
