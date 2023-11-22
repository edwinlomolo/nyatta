package model

type MobileMoneyPayload struct {
	Phone    string `json:"phone"`
	Provider string `json:"provider"`
}

type PaystackMpesaChargePayload struct {
	Amount      int                `json:"amount"`
	Email       string             `json:"email"`
	Currency    string             `json:"currency"`
	MobileMoney MobileMoneyPayload `json:"mobile_money"`
}

type PaystackMpesaChargeResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Reference   string `json:"reference"`
		Status      string `json:"status"`
		DisplayText string `json:"display_text"`
	} `json:"data"`
}

type PaystackAuthorization struct {
	Bank        string `json:"bank"`
	Channel     string `json:"channel"`
	CountryCode string `json:"country_code"`
	Brand       string `json:"brand"`
	AuthCode    string `json:"authorization_code"`
}

type Customer struct {
	ID           int    `json:"id"`
	Phone        string `json:"phone"`
	CustomerCode string `json:"customer_code"`
	Email        string `json:"email"`
}

type CallbackData struct {
	Status        string                `json:"success"`
	Reference     string                `json:"reference"`
	Amount        int                   `json:"amount"`
	PaidAt        string                `json:"paid_at"`
	Customer      Customer              `json:"customer"`
	CreatedAt     string                `json:"created_at"`
	Channel       string                `json:"channel"`
	Currency      string                `json:"currency"`
	Fees          int                   `json:"fees"`
	Authorization PaystackAuthorization `json:"authorization"`
}

type PaystackCallbackResponse struct {
	Event string       `json:"event"`
	Data  CallbackData `json:"data"`
}
