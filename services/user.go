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
			FirstName:  foundUser.FirstName.String,
			LastName:   foundUser.LastName.String,
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
		FirstName:  foundUser.FirstName.String,
		LastName:   foundUser.LastName.String,
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

// UpdateUserInfo - update user details
func (u *UserServices) UpdateUserInfo(userId int64, phone, firstName, lastName, avatar string) (*model.User, error) {
	foundUpload, err := u.GetUserAvatar(userId)
	if err != nil {
		u.log.Errorf("%s:%v", u.ServiceName(), err)
		return nil, err
	}

	if foundUpload == nil {
		if _, err := u.queries.CreateUserAvatar(ctx, sqlStore.CreateUserAvatarParams{
			Upload:   avatar,
			Category: model.UploadCategoryProfileImg.String(),
			UserID:   sql.NullInt64{Int64: userId, Valid: true},
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
		return &model.User{ID: strconv.FormatInt(updatedUser.ID, 10)}, nil
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
		uploadId, err := strconv.ParseInt(foundUpload.ID, 10, 64)
		if err != nil {
			u.log.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
		if _, err := u.queries.UpdateUpload(ctx, sqlStore.UpdateUploadParams{
			ID:     uploadId,
			Upload: avatar,
		}); err != nil {
			u.log.Errorf("%s:%v", u.ServiceName(), err)
			return nil, err
		}
		return &model.User{ID: strconv.FormatInt(updatedUser.ID, 10)}, nil
	}
}

// GetUserAvatar - grab avatar
func (u *UserServices) GetUserAvatar(userId int64) (*model.AnyUpload, error) {
	foundUpload, err := u.queries.GetUserAvatar(ctx, sqlStore.GetUserAvatarParams{
		UserID:   sql.NullInt64{Int64: userId, Valid: true},
		Category: model.UploadCategoryProfileImg.String(),
	})
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	return &model.AnyUpload{
		ID:     strconv.FormatInt(foundUpload.ID, 10),
		Upload: foundUpload.Upload,
	}, nil
}
