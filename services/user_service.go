package services

import (
	"fmt"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *model.NewUser) (*model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	SignIn(user *model.NewUser) (*string, error)
	ValidateToken(token *string) (*jwt.Token, error)
}

type UserServices struct {
	store *gorm.DB
	log   *zap.SugaredLogger
	auth  *AuthServices
}

var _ UserService = &UserServices{}

func NewUserService(store *gorm.DB, logger *zap.SugaredLogger, config *nyatta_context.Config) *UserServices {
	authServices := NewAuthService(logger, config)
	return &UserServices{store, logger, authServices}
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

func (u *UserServices) SignIn(user *model.NewUser) (*string, error) {
	// Find user returning user
	var newUser *model.User
	var err error
	newUser, err = u.FindByEmail(user.Email)
	if err != nil {
		return nil, fmt.Errorf("Error signing in user: %v", err)
	}
	// User cannot be found - new user
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

func (u *UserServices) FindById(id string) (*model.User, error) {
	var foundUser *model.User
	if err := u.store.Where("id = ?", id).Preload("Properties").Find(&foundUser).Error; err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (u *UserServices) FindByEmail(email string) (*model.User, error) {
	var foundUser *model.User
	if err := u.store.Where("email = ?", email).Preload("Properties").Find(&foundUser).Error; err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (u *UserServices) ValidateToken(tokenString *string) (*jwt.Token, error) {
	token, err := u.auth.ValidateJWT(tokenString)
	return token, err
}

func (u *UserServices) ServiceName() string {
	return "UserServices"
}
