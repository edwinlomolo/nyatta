package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
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
	queries *sqlStore.Queries
	log     *log.Logger
	auth    *AuthServices
}

// _ - UserServices{} implements UserService
var _ UserService = &UserServices{}

func NewUserService(queries *sqlStore.Queries, logger *log.Logger, config *config.Jwt) *UserServices {
	authServices := NewAuthService(logger, config)
	return &UserServices{queries, logger, authServices}
}

// CreateUser - create a new user
func (u *UserServices) CreateUser(user *model.NewUser) (*model.User, error) {
	ctx := context.Background()
	insertedUser, err := u.queries.CreateUser(ctx, sqlStore.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        strconv.FormatInt(insertedUser.ID, 10),
		FirstName: insertedUser.FirstName,
		LastName:  insertedUser.LastName,
		Email:     insertedUser.Email,
	}, nil
}

// TODO
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
	ctx := context.Background()
	propertyId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	foundUser, err := u.queries.GetUser(ctx, int64(propertyId))
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	return &model.User{
		ID:        strconv.FormatInt(foundUser.ID, 10),
		FirstName: foundUser.FirstName,
		LastName:  foundUser.LastName,
		Email:     foundUser.Email,
	}, nil
}

// TODO
// FindByEmail - return user given user email
func (u *UserServices) FindByEmail(email string) (*model.User, error) {
	return nil, nil
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
