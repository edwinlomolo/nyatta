package services

import (
	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

// TwilioService - represent twilio services
type TwilioService interface {
	SendVerification(phone string) (string, error)
	VerifyCode(phone, verifyCode string) (string, error)
}

type twilioClient struct {
	Client      *twilio.RestClient
	Sid         string // Verify service id
	userService UserService
	queries     *sqlStore.Queries
	logger      *logrus.Logger
}

// NewTwilioService - create new instance of Twilio services
func NewTwilioService(cfg config.TwilioConfig, queries *sqlStore.Queries, logger *logrus.Logger) TwilioService {
	// Create twilio client
	twilio := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.Sid,
		Password: cfg.AuthToken,
	})

	return &twilioClient{Client: twilio, Sid: cfg.VerifySid, queries: queries, logger: logger}
}

// SendVerification - sends verification code
func (t twilioClient) SendVerification(phone string) (string, error) {
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
func (t twilioClient) VerifyCode(phone, verifyCode string) (string, error) {
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
func (t *twilioClient) ServiceName() string {
	return "twilioClient"
}
