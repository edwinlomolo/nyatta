package config

type TwilioConfig struct {
	Sid       string `json:"sid"`
	AuthToken string `json:"authToken"`
	VerifySid string `json:"verifySid"`
}
