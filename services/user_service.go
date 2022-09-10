package services

import (
	"fmt"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
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
	// User exists
	var existingUser *model.User
	if err := u.store.Where("email = ?", user.Email).Find(&existingUser).Error; err != nil {
		return nil, fmt.Errorf("%s: %v", nyatta_context.DatabaseError, err)
	}
	if existingUser.Email != "" {
		u.log.Errorf("User with email %s already exists. Please login", user.Email)
		return nil, nyatta_context.AlreadyExists
	}
	newUser := &model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	if err := u.store.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *UserServices) GetUser(id string) (*model.User, error) {
	var foundUser *model.User
	// Id can be actual user id/email
	if err := u.store.Where("id = ?", id).Find(&foundUser).Error; err != nil {
		return nil, err
	}

	if foundUser.ID == "" {
		if err := u.store.Where("email = ?", id).Find(&foundUser).Error; err != nil {
			return nil, err
		}
	}
	return foundUser, nil
}

func (u *UserServices) ServiceName() string {
	return "UserServices"
}
