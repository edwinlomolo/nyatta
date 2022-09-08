package services

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *model.NewUser) (*model.User, error)
	GetUser(id string) (*model.User, error)
}

type UserServices struct {
	store *gorm.DB
	log   *zap.SugaredLogger
}

func NewUserService(store *gorm.DB, logger *zap.SugaredLogger) *UserServices {
	return &UserServices{store, logger}
}

func (u *UserServices) CreateUser(user *model.NewUser) (*model.User, error) {
	return &model.User{}, nil
}

func (u *UserServices) GetUser(id string) (*model.User, error) {
	return &model.User{}, nil
}

func (u *UserServices) ServiceName() string {
	return "UserServices"
}
