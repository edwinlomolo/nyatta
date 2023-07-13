package services

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/nyaruka/phonenumbers"
	"github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioServices struct {
	Client      *twilio.RestClient
	Sid         string // Verify service id
	userService *UserServices
	queries     *sqlStore.Queries
	logger      *logrus.Logger
}

// TwilioServices implements Twilio
var _ interfaces.Twilio = &TwilioServices{}

// NewTwilioService - create new instance of Twilio services
func NewTwilioService(cfg config.TwilioConfig, queries *sqlStore.Queries, logger *logrus.Logger) *TwilioServices {
	// Create twilio client
	twilio := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.Sid,
		Password: cfg.AuthToken,
	})

	return &TwilioServices{Client: twilio, Sid: cfg.VerifySid, queries: queries, logger: logger}
}

// SendVerification - sends verification code
func (t TwilioServices) SendVerification(phone string, countryCode model.CountryCode) (string, error) {
	num, err := phonenumbers.Parse(phone, countryCode.String())
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", err
	}

	formattedNumber := phonenumbers.Format(num, phonenumbers.INTERNATIONAL) // Get number international format
	params := &verify.CreateVerificationParams{}
	params.SetTo(formattedNumber)
	params.SetChannel("sms")

	res, err := t.Client.VerifyV2.CreateVerification(t.Sid, params)
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", err
	} else {
		if res.Status != nil {
			return *res.Status, nil
		}
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", errors.New("nil response from twilio")
	}
}

// VerifyCode - verify verification code
func (t TwilioServices) VerifyCode(phone, verifyCode string, countryCode model.CountryCode) (string, error) {
	num, err := phonenumbers.Parse(phone, countryCode.String())
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", err
	}

	formattedNumber := phonenumbers.Format(num, phonenumbers.INTERNATIONAL) // Get number international format
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(formattedNumber)
	params.SetCode(verifyCode)

	res, err := t.Client.VerifyV2.CreateVerificationCheck(t.Sid, params)
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", err
	} else {
		if res.Status != nil {
			return *res.Status, nil
		}
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", errors.New("nil response from twilio")
	}
}

// UpdateUserPhone - update user phone number
func (t *TwilioServices) UpdateUserPhone(email, phone string) (*model.User, error) {
	updatedUser, err := t.queries.UpdateUserPhone(ctx, sqlStore.UpdateUserPhoneParams{
		Email: sql.NullString{String: email, Valid: true},
		Phone: sql.NullString{String: phone, Valid: true},
	})
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return nil, err
	}
	return &model.User{
		ID:         strconv.FormatInt(updatedUser.ID, 10),
		FirstName:  updatedUser.FirstName.String,
		LastName:   updatedUser.LastName.String,
		Email:      updatedUser.Email.String,
		Phone:      updatedUser.Phone.String,
		Onboarding: updatedUser.Onboarding.Bool,
		CreatedAt:  &updatedUser.CreatedAt,
		UpdatedAt:  &updatedUser.UpdatedAt,
	}, nil
}

// ServiceName - returns service name
func (t TwilioServices) ServiceName() string {
	return "TwilioServices"
}
