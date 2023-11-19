package services

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// UserServices - represents user service
type UserServices struct {
	queries   *sqlStore.Queries
	log       *logrus.Logger
	auth      *AuthServices
	twilio    *TwilioServices
	sendEmail SendEmail
	env       string
}

// _ - UserServices{} implements UserService
var _ interfaces.UserService = &UserServices{}

func NewUserService(queries *sqlStore.Queries, logger *logrus.Logger, env string, config *config.Jwt, twilio *TwilioServices, sendEmail SendEmail) *UserServices {
	authServices := NewAuthService(logger, config)
	return &UserServices{queries, logger, authServices, twilio, sendEmail, env}
}

// FindUserByPhone - get user by phone number
func (u *UserServices) FindUserByPhone(phone string) (*model.User, error) {
	var foundUser sqlStore.User
	var err error

	foundUser, err = u.queries.FindUserByPhone(ctx, phone)
	if err != nil && err == sql.ErrNoRows {
		foundUser, err = u.queries.CreateUser(ctx, phone)
		if err != nil {
			u.log.Errorf("%s: %v", u.ServiceName(), err)
			return nil, err
		}

		isLandlord := time.Now().Before(foundUser.NextRenewal)
		return &model.User{
			ID:         strconv.FormatInt(foundUser.ID, 10),
			IsLandlord: isLandlord,
			Phone:      foundUser.Phone,
			CreatedAt:  &foundUser.CreatedAt,
			UpdatedAt:  &foundUser.UpdatedAt,
		}, nil
	} else if err != nil && err != sql.ErrNoRows {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}

	isLandlord := time.Now().Before(foundUser.NextRenewal)
	return &model.User{
		ID:         strconv.FormatInt(foundUser.ID, 10),
		Phone:      foundUser.Phone,
		IsLandlord: isLandlord,
		CreatedAt:  &foundUser.CreatedAt,
		UpdatedAt:  &foundUser.UpdatedAt,
	}, nil
}

// SignIn - signin existing/returning user
func (u *UserServices) SignIn(user *model.NewUser) (*model.SignIn, error) {
	signInResponse := &model.SignIn{}

	// user - existing user
	var newUser *model.User
	var err error
	newUser, err = u.FindUserByPhone(user.Phone)
	if err != nil {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}
	token, err := u.auth.SignJWT(newUser)
	if err != nil {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}

	signInResponse.Token = *token
	signInResponse.User = newUser

	return signInResponse, nil
}

// ValidateToken - validate jwt token
func (u *UserServices) ValidateToken(tokenString *string) (*jwt.Token, error) {
	token, err := u.auth.ValidateJWT(tokenString)
	return token, err
}

// ServiceName - return service name
func (u UserServices) ServiceName() string {
	return "UserServices"
}
