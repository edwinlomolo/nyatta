package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database/store"
	"github.com/sirupsen/logrus"
)

type PaystackServices struct {
	config   config.Paystack
	logger   *logrus.Logger
	baseApi  string
	sqlStore *store.Queries
}

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

func NewPaystackService(cfg config.Paystack, logger *logrus.Logger, sqlStore *store.Queries) *PaystackServices {
	return &PaystackServices{config: cfg, logger: logger, baseApi: cfg.BaseApi, sqlStore: sqlStore}
}

func (p *PaystackServices) ChargeMpesaPhone(payload PaystackMpesaChargePayload) (*PaystackMpesaChargeResponse, error) {
	var chargeResponse *PaystackMpesaChargeResponse

	url := p.baseApi + "/charge"
	payload.MobileMoney.Provider = "mpesa"
	if p.config.Env == "test" {
		payload.MobileMoney.Phone = "+254710000000"
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+p.config.SecretKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&chargeResponse); err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	return chargeResponse, nil
}

func (p PaystackServices) ServiceName() string {
	return "PaystackServices"
}
