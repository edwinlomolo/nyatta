package services

import (
	"github.com/3dw1nM0535/nyatta/interfaces"
)

type TenancyServices struct{}

// _ - TenancyServices{} implements TenancyService interface
var _ interfaces.TenancyService = &TenancyServices{}

// NewTenancyService - factory for tenancy services
func NewTenancyService() *TenancyServices {
	return &TenancyServices{}
}

// ServiceName - return service name
func (t TenancyServices) ServiceName() string {
	return "TenancyServices"
}
