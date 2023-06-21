package services

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	jwt "github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

// UserServices - represents user service
type UserServices struct {
	queries *sqlStore.Queries
	log     *log.Logger
	auth    *AuthServices
}

// _ - UserServices{} implements UserService
var _ interfaces.UserService = &UserServices{}

func NewUserService(queries *sqlStore.Queries, logger *log.Logger, config *config.Jwt) *UserServices {
	authServices := NewAuthService(logger, config)
	return &UserServices{queries, logger, authServices}
}

// CreateUser - create a new user
func (u *UserServices) CreateUser(user *model.NewUser) (*model.User, error) {
	ctx := context.Background()
	insertedUser, err := u.queries.CreateUser(ctx, sqlStore.CreateUserParams{
		FirstName: sql.NullString{String: user.FirstName},
		LastName:  sql.NullString{String: user.LastName},
		Email:     sql.NullString{String: user.Email},
		Avatar:    sql.NullString{String: user.Avatar},
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:         strconv.FormatInt(insertedUser.ID, 10),
		FirstName:  insertedUser.FirstName.String,
		LastName:   insertedUser.LastName.String,
		Email:      insertedUser.Email.String,
		Onboarding: insertedUser.Onboarding.Bool,
		Avatar:     insertedUser.Avatar.String,
		CreatedAt:  &insertedUser.CreatedAt,
		UpdatedAt:  &insertedUser.UpdatedAt,
	}, nil
}

// SignIn - signin existing/returning user
func (u *UserServices) SignIn(user *model.NewUser) (*string, error) {
	// user - existing user
	var newUser *model.User
	var err error
	newUser, err = u.FindByEmail(user.Email)
	if err != nil && err.Error() != "User not found" {
		return nil, err
	}
	// user - new user
	if newUser == nil {
		newUser, err = u.CreateUser(user)
		if err != nil {
			return nil, err
		}
	}
	token, err := u.auth.SignJWT(newUser)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// FindById - return user given user id
func (u *UserServices) FindById(id string) (*model.User, error) {
	ctx := context.Background()
	propertyId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	foundUser, err := u.queries.GetUser(ctx, propertyId)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	return &model.User{
		ID:        strconv.FormatInt(foundUser.ID, 10),
		FirstName: foundUser.FirstName.String,
		LastName:  foundUser.LastName.String,
		Email:     foundUser.Email.String,
		Avatar:    foundUser.Avatar.String,
		CreatedAt: &foundUser.CreatedAt,
		UpdatedAt: &foundUser.UpdatedAt,
	}, nil
}

// FindByEmail - return user given user email
func (u *UserServices) FindByEmail(email string) (*model.User, error) {
	ctx := context.Background()
	foundUser, err := u.queries.FindByEmail(ctx, sql.NullString{String: email})
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	return &model.User{
		ID:        strconv.FormatInt(foundUser.ID, 10),
		FirstName: foundUser.FirstName.String,
		LastName:  foundUser.LastName.String,
		Email:     foundUser.Email.String,
		Avatar:    foundUser.Avatar.String,
		CreatedAt: &foundUser.CreatedAt,
		UpdatedAt: &foundUser.UpdatedAt,
	}, nil
}

// UpdateUser - update user details
func (u *UserServices) UpdateUser(input *model.UpdateUserInput) (*model.User, error) {
	userId, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	// TODO any other way around casting to NullString?
	updatedUser, err := u.queries.UpdateUser(ctx, sqlStore.UpdateUserParams{
		ID:         userId,
		FirstName:  sql.NullString{String: input.FirstName, Valid: true},
		LastName:   sql.NullString{String: input.LastName, Valid: true},
		Avatar:     sql.NullString{String: input.Avatar, Valid: true},
		Onboarding: sql.NullBool{Bool: input.Onboarding, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	log.Println(updatedUser.Onboarding.Bool)
	return &model.User{
		ID:         strconv.FormatInt(updatedUser.ID, 10),
		FirstName:  updatedUser.FirstName.String,
		LastName:   updatedUser.LastName.String,
		Email:      updatedUser.Email.String,
		Onboarding: updatedUser.Onboarding.Bool,
		Avatar:     updatedUser.Avatar.String,
		CreatedAt:  &updatedUser.CreatedAt,
		UpdatedAt:  &updatedUser.UpdatedAt,
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

// FindUserByPhone - get user by phone number
func (u *UserServices) FindUserByPhone(phone string) (*model.User, error) {
	phoneNumber := sql.NullString{String: phone, Valid: true}
	var foundUser sqlStore.User
	var err error
	foundUser, err = u.queries.FindUserByPhone(ctx, phoneNumber)
	if err == sql.ErrNoRows {
		// Create new user(auto-onboard)
		foundUser, err = u.queries.CreateUser(ctx, sqlStore.CreateUserParams{
			Phone: phoneNumber,
		})
		if err != nil {
			return nil, err
		}
		return &model.User{
			ID:         strconv.FormatInt(foundUser.ID, 10),
			Onboarding: foundUser.Onboarding.Bool,
		}, nil
	} else if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &model.User{
		ID:         strconv.FormatInt(foundUser.ID, 10),
		FirstName:  foundUser.FirstName.String,
		LastName:   foundUser.LastName.String,
		Email:      foundUser.Email.String,
		Onboarding: foundUser.Onboarding.Bool,
		Avatar:     foundUser.Avatar.String,
		CreatedAt:  &foundUser.CreatedAt,
		UpdatedAt:  &foundUser.UpdatedAt,
	}, nil
}
