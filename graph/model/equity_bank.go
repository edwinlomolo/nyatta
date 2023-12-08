package model

type AuthRequest struct {
	MerchantCode   string `json:"merchantCode"`
	ConsumerSecret string `json:"consumerSecret"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    string `json:"expiresIn"`
	IssuedAt     string `json:"issuedAt"`
}
