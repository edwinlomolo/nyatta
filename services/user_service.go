package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserServices struct {
	store *gorm.DB
	log   *zap.Logger
}

func NewUserService(store *gorm.DB, logger *zap.Logger) *UserServices {
	return &UserServices{store, logger}
}
