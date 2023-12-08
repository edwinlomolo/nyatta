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

type Balance struct {
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type AccountBalanceResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Balances []*Balance `json:"balances"`
	} `json:"data"`
}
