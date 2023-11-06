package services

import (
	"database/sql"
	"errors"
	"os"
	"strconv"

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
		// Create new user(auto-onboard)
		foundUser, err = u.queries.CreateUser(ctx, phone)
		if err != nil {
			u.log.Errorf("%s: %v", u.ServiceName(), err)
			return nil, err
		}
		return &model.User{
			ID:         strconv.FormatInt(foundUser.ID, 10),
			IsLandlord: foundUser.IsLandlord.Bool,
			Phone:      foundUser.Phone,
			CreatedAt:  &foundUser.CreatedAt,
		}, nil
	} else if err != nil && err != sql.ErrNoRows {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}
	return &model.User{
		ID:         strconv.FormatInt(foundUser.ID, 10),
		Phone:      foundUser.Phone,
		IsLandlord: foundUser.IsLandlord.Bool,
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

// FindById - return user given user id
func (u *UserServices) FindById(id string) (*model.User, error) {
	propertyId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}
	foundUser, err := u.queries.GetUser(ctx, propertyId)
	if err == sql.ErrNoRows {
		u.log.Errorf("%s: %v", u.ServiceName(), errors.New("User not found"))
		return nil, errors.New("User not found")
	}
	return &model.User{
		ID:        strconv.FormatInt(foundUser.ID, 10),
		FirstName: foundUser.FirstName.String,
		LastName:  foundUser.LastName.String,
		Email:     foundUser.Email.String,
		CreatedAt: &foundUser.CreatedAt,
		UpdatedAt: &foundUser.UpdatedAt,
	}, nil
}

// FindByEmail - return user given user email
func (u *UserServices) FindByEmail(email string) (*model.User, error) {
	foundUser, err := u.queries.FindByEmail(ctx, sql.NullString{String: email, Valid: true})
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	return &model.User{
		ID:         strconv.FormatInt(foundUser.ID, 10),
		FirstName:  foundUser.FirstName.String,
		LastName:   foundUser.LastName.String,
		Email:      foundUser.Email.String,
		Onboarding: foundUser.Onboarding.Bool,
		Phone:      foundUser.Phone,
		CreatedAt:  &foundUser.CreatedAt,
		UpdatedAt:  &foundUser.UpdatedAt,
	}, nil
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

// OnboardUser - update user onboarding status
func (u *UserServices) OnboardUser(email string, onboarding bool) (*model.User, error) {
	onboardedUser, err := u.queries.OnboardUser(ctx, sqlStore.OnboardUserParams{
		Email:      sql.NullString{String: email, Valid: true},
		Onboarding: sql.NullBool{Bool: onboarding, Valid: true},
	})
	if err != nil {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}
	// send welcome email in staging/prod env
	if u.env == "staging" || u.env == "production" {
		from := os.Getenv("EMAIL_FROM")
		err := u.sendEmail([]string{onboardedUser.Email.String}, from, "Welcome to Nyattta", newUserEmail)
		if err != nil {
			u.log.Errorf("Error sending email:%s: %v", u.ServiceName(), err)
			return nil, err
		}
	}
	return &model.User{
		ID:         strconv.FormatInt(onboardedUser.ID, 10),
		FirstName:  onboardedUser.FirstName.String,
		LastName:   onboardedUser.LastName.String,
		Email:      onboardedUser.Email.String,
		Onboarding: onboardedUser.Onboarding.Bool,
		Phone:      onboardedUser.Phone,
		CreatedAt:  &onboardedUser.CreatedAt,
		UpdatedAt:  &onboardedUser.UpdatedAt,
	}, nil
}
