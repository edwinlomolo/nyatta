package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
)

type Twilio interface {
	SendVerification(phone string, countryCode model.CountryCode) (string, error)
	VerifyCode(phone, email, verifyCode string, countryCode model.CountryCode) (string, error)
}
