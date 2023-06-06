package services

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/nyaruka/phonenumbers"
	log "github.com/sirupsen/logrus"
)

type TwilioServices struct{}

// TwilioServices implements Twilio
var _ interfaces.Twilio = &TwilioServices{}

// NewTwilioService - create new instance of Twilio services
func NewTwilioService() *TwilioServices {
	return &TwilioServices{}
}

// SendVerification - sends verification code
func (t TwilioServices) SendVerification(phone string, countryCode model.CountryCode) (string, error) {
	num, err := phonenumbers.Parse(phone, countryCode.String())
	if err != nil {
		return "", err
	}
	log.Infoln(num)
	return "", nil
}

// VerifyCode - verify verification code
func (t TwilioServices) VerifyCode(phone, verifyCode string, countryCode model.CountryCode) (string, error) {
	_, err := phonenumbers.Parse(phone, countryCode.String())
	if err != nil {
		return "", err
	}
	log.Infoln(verifyCode)
	return "", nil
}
