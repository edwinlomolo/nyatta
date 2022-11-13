package services

import (
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *model.NewUser) (*model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	SignIn(user *model.NewUser) (*string, error)
	ValidateToken(token *string) (*jwt.Token, error)
}

// UserServices - represents user service
type UserServices struct {
	store *gorm.DB
	log   *log.Logger
	auth  *AuthServices
}

// _ - UserServices{} implements UserService
var _ UserService = &UserServices{}

func NewUserService(store *gorm.DB, logger *log.Logger, config *config.Jwt) *UserServices {
	authServices := NewAuthService(logger, config)
	return &UserServices{store, logger, authServices}
}

// CreateUser - create a new user
func (u *UserServices) CreateUser(user *model.NewUser) (*model.User, error) {
	// User exists
	var existingUser *model.User
	if err := u.store.Where("email = ?", user.Email).Find(&existingUser).Error; err != nil {
		return nil, fmt.Errorf("%s: %v", config.DatabaseError, err)
	}
	if existingUser.Email != "" {
		u.log.Errorf("User with email %s already exists. Please login", user.Email)
		return nil, config.AlreadyExists
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

// SignIn - signin existing/returning user
func (u *UserServices) SignIn(user *model.NewUser) (*string, error) {
	// user - existing user
	var newUser *model.User
	var err error
	newUser, err = u.FindByEmail(user.Email)
	if err != nil {
		return nil, fmt.Errorf("Error signing in user: %v", err)
	}
	// user - new user
	if newUser == nil {
		newUser, err = u.CreateUser(user)
		if err != nil {
			return nil, fmt.Errorf("Error creating new user: %v", err)
		}
	}
	token, err := u.auth.SignJWT(newUser)
	if err != nil {
		return nil, fmt.Errorf("Error signing user token: %v", err)
	}
	return token, nil
}

// FindById - return user given user id
func (u *UserServices) FindById(id string) (*model.User, error) {
	var foundUser *model.User
	if err := u.store.Where("id = ?", id).Preload("Properties").Find(&foundUser).Error; err != nil {
		return nil, err
	}
	return foundUser, nil
}

// FindByEmail - return user given user email
func (u *UserServices) FindByEmail(email string) (*model.User, error) {
	var foundUser *model.User
	if err := u.store.Where("email = ?", email).Preload("Properties").Find(&foundUser).Error; err != nil {
		return nil, err
	}
	return foundUser, nil
}

// ValidateToken - validate jwt token
func (u *UserServices) ValidateToken(tokenString *string) (*jwt.Token, error) {
	token, err := u.auth.ValidateJWT(tokenString)
	return token, err
}

// ServiceName - return service name
func (u *UserServices) ServiceName() string {
	return "UserServices"
}
