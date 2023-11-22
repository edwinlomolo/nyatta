package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UserServices - represents user service
type UserServices struct {
	queries *sqlStore.Queries
	log     *logrus.Logger
	auth    *AuthServices
	twilio  *TwilioServices
}

// _ - UserServices{} implements UserService
var _ interfaces.UserService = &UserServices{}

func NewUserService(queries *sqlStore.Queries, logger *logrus.Logger, config *config.Jwt, twilio *TwilioServices) *UserServices {
	authServices := NewAuthService(logger, config)
	return &UserServices{queries, logger, authServices, twilio}
}

// FindUserByPhone - get user by phone number
func (u *UserServices) FindUserByPhone(ctx context.Context, phone string) (*model.User, error) {
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
			ID:               foundUser.ID,
			IsLandlord:       isLandlord,
			FirstName:        &foundUser.FirstName.String,
			LastName:         &foundUser.LastName.String,
			Phone:            foundUser.Phone,
			SubscribeRetries: int(foundUser.SubscribeRetries),
			CreatedAt:        &foundUser.CreatedAt,
			UpdatedAt:        &foundUser.UpdatedAt,
		}, nil
	} else if err != nil && err != sql.ErrNoRows {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}

	isLandlord := time.Now().Before(foundUser.NextRenewal)
	return &model.User{
		ID:               foundUser.ID,
		Phone:            foundUser.Phone,
		FirstName:        &foundUser.FirstName.String,
		LastName:         &foundUser.LastName.String,
		IsLandlord:       isLandlord,
		SubscribeRetries: int(foundUser.SubscribeRetries),
		CreatedAt:        &foundUser.CreatedAt,
		UpdatedAt:        &foundUser.UpdatedAt,
	}, nil
}

// SignIn - signin existing/returning user
func (u *UserServices) SignIn(ctx context.Context, user *model.NewUser) (*model.SignIn, error) {
	signInResponse := &model.SignIn{}

	// user - existing user
	var newUser *model.User
	var err error
	newUser, err = u.FindUserByPhone(ctx, user.Phone)
	if err != nil {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}
	token, err := u.auth.SignJWT(ctx, newUser)
	if err != nil {
		u.log.Errorf("%s: %v", u.ServiceName(), err)
		return nil, err
	}

	signInResponse.Token = *token
	signInResponse.User = newUser

	return signInResponse, nil
}

// ValidateToken - validate jwt token
func (u *UserServices) ValidateToken(ctx context.Context, tokenString *string) (*jwt.Token, error) {
	token, err := u.auth.ValidateJWT(ctx, tokenString)
	return token, err
}

// ServiceName - return service name
func (u UserServices) ServiceName() string {
	return "UserServices"
}

// UpdateUserInfo - update user details
func (u *UserServices) UpdateUserInfo(ctx context.Context, userId uuid.UUID, firstName, lastName, avatar string) (*model.User, error) {
	phone := ctx.Value("phone").(string)

	foundUpload, err := u.GetUserAvatar(ctx, userId)
	if err != nil {
		u.log.Errorf("%s:%v", u.ServiceName(), err)
		return nil, err
	}

	if foundUpload == nil {
		if _, err := u.queries.CreateUserAvatar(ctx, sqlStore.CreateUserAvatarParams{
			Upload:   avatar,
			Category: model.UploadCategoryProfileImg.String(),
			UserID:   uuid.NullUUID{UUID: userId, Valid: true},
		}); err != nil {
			u.log.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
		updatedUser, err := u.queries.UpdateUserInfo(ctx, sqlStore.UpdateUserInfoParams{
			FirstName: sql.NullString{String: firstName, Valid: true},
			LastName:  sql.NullString{String: lastName, Valid: true},
			Phone:     phone,
		})
		if err != nil {
			u.log.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
		return &model.User{ID: updatedUser.ID}, nil
	} else {
		updatedUser, err := u.queries.UpdateUserInfo(ctx, sqlStore.UpdateUserInfoParams{
			FirstName: sql.NullString{String: firstName, Valid: true},
			LastName:  sql.NullString{String: lastName, Valid: true},
			Phone:     phone,
		})
		if err != nil {
			u.log.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}

		if _, err := u.queries.UpdateUpload(ctx, sqlStore.UpdateUploadParams{
			ID:     foundUpload.ID,
			Upload: avatar,
		}); err != nil {
			u.log.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
		return &model.User{ID: updatedUser.ID}, nil
	}
}

// GetUserAvatar - grab avatar
func (u *UserServices) GetUserAvatar(ctx context.Context, userId uuid.UUID) (*model.AnyUpload, error) {
	foundUpload, err := u.queries.GetUserAvatar(ctx, sqlStore.GetUserAvatarParams{
		UserID:   uuid.NullUUID{UUID: userId, Valid: true},
		Category: model.UploadCategoryProfileImg.String(),
	})
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	return &model.AnyUpload{
		ID:     foundUpload.ID,
		Upload: foundUpload.Upload,
	}, nil
}

// GetUser - grab user
func (u *UserServices) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	foundUser, err := u.queries.GetUser(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	isLandlord := time.Now().Before(foundUser.NextRenewal)
	return &model.User{
		ID:               foundUser.ID,
		FirstName:        &foundUser.FirstName.String,
		LastName:         &foundUser.LastName.String,
		IsLandlord:       isLandlord,
		SubscribeRetries: int(foundUser.SubscribeRetries),
		Phone:            foundUser.Phone,
		CreatedAt:        &foundUser.CreatedAt,
		UpdatedAt:        &foundUser.UpdatedAt,
	}, nil
}
