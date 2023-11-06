package interfaces

type Twilio interface {
	SendVerification(phone string) (string, error)
	VerifyCode(phone, verifyCode string) (string, error)
}
