package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PropertyService interface {
	ServiceName() string
}

type PropertyServices struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewPropertyService(store *gorm.DB, logger *zap.SugaredLogger) *PropertyServices {
	return &PropertyServices{db: store, logger: logger}
}

func (p PropertyServices) ServiceName() string {
	return "PropertyServices"
}
