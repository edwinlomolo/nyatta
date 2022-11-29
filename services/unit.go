package services

import (
	"github.com/3dw1nM0535/nyatta/interfaces"
)

type UnitServices struct{}

// _ - UnitServices{} implements UnitService interface
var _ interfaces.UnitService = &UnitServices{}

// NewUnitService - factory for UnitServices
func NewUnitService() *UnitServices {
	return &UnitServices{}
}

// ServiceName - return service name
func (u UnitServices) ServiceName() string {
	return "UnitServices"
}
