package services

import (
	"errors"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/nyaruka/phonenumbers"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioServices struct {
	Client      *twilio.RestClient
	Sid         string // Verify service id
	userService *UserServices
}

// TwilioServices implements Twilio
var _ interfaces.Twilio = &TwilioServices{}

// NewTwilioService - create new instance of Twilio services
func NewTwilioService(cfg config.TwilioConfig, userServices *UserServices) *TwilioServices {
	// Create twilio client
	twilio := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.Sid,
		Password: cfg.AuthToken,
	})

	return &TwilioServices{Client: twilio, Sid: cfg.VerifySid, userService: userServices}
}

// SendVerification - sends verification code
func (t TwilioServices) SendVerification(phone string, countryCode model.CountryCode) (string, error) {
	num, err := phonenumbers.Parse(phone, countryCode.String())
	if err != nil {
		return "", err
	}

	formattedNumber := phonenumbers.Format(num, phonenumbers.INTERNATIONAL) // Get number international format
	params := &verify.CreateVerificationParams{}
	params.SetTo(formattedNumber)
	params.SetChannel("sms")

	res, err := t.Client.VerifyV2.CreateVerification(t.Sid, params)
	if err != nil {
		return "", err
	} else {
		if res.Status != nil {
			return *res.Status, nil
		}
		return "", errors.New("nil response from twilio")
	}
}

// VerifyCode - verify verification code
func (t TwilioServices) VerifyCode(phone, email, verifyCode string, countryCode model.CountryCode) (string, error) {
	num, err := phonenumbers.Parse(phone, countryCode.String())
	if err != nil {
		return "", err
	}

	formattedNumber := phonenumbers.Format(num, phonenumbers.INTERNATIONAL) // Get number international format
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(formattedNumber)
	params.SetCode(verifyCode)

	res, err := t.Client.VerifyV2.CreateVerificationCheck(t.Sid, params)
	if err != nil {
		return "", err
	} else {
		if res.Status != nil {
			if *res.Status == "approved" {
				// Start creating user after verification
				_, err := t.userService.UpdateUserPhone(email, phone)
				if err != nil {
					return "", err
				}
			}
			return *res.Status, nil
		}
		return "", errors.New("nil response from twilio")
	}
}
