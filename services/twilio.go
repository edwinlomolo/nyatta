package services

import (
	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/interfaces"
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
func (t TwilioServices) SendVerification(phone string) (string, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	res, err := t.Client.VerifyV2.CreateVerification(t.Sid, params)
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", err
	} else {
		if res.Status != nil {
			return *res.Status, nil
		}
		t.logger.Errorf("%s: %v", t.ServiceName(), config.TwilioNilErr)
		return "", config.TwilioNilErr
	}
}

// VerifyCode - verify verification code
func (t TwilioServices) VerifyCode(phone, verifyCode string) (string, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(verifyCode)

	res, err := t.Client.VerifyV2.CreateVerificationCheck(t.Sid, params)
	if err != nil {
		t.logger.Errorf("%s: %v", t.ServiceName(), err)
		return "", err
	} else {
		if res.Status != nil {
			return *res.Status, nil
		}
		t.logger.Errorf("%s: %v", t.ServiceName(), config.TwilioNilErr)
		return "", config.TwilioNilErr
	}
}

// ServiceName - returns service name
func (t TwilioServices) ServiceName() string {
	return "TwilioServices"
}
