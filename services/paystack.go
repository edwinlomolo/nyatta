package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
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

type PaystackAuthorization struct {
	Bank              string `json:"bank"`
	Channel           string `json:"channel"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	MobileMoneyNumber string `json:"mobile_money_number"`
	AuthCode          string `json:"authorization_code"`
}

type CallbackData struct {
	Status        string                `json:"success"`
	Reference     string                `json:"reference"`
	Amount        int                   `json:"amount"`
	PaidAt        string                `json:"paid_at"`
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

func NewPaystackService(cfg config.Paystack, logger *logrus.Logger, sqlStore *store.Queries) *PaystackServices {
	return &PaystackServices{config: cfg, logger: logger, baseApi: cfg.BaseApi, sqlStore: sqlStore}
}

func (p *PaystackServices) ChargeMpesaPhone(phone string, payload PaystackMpesaChargePayload) (*PaystackMpesaChargeResponse, error) {
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

	if _, err := p.sqlStore.CreateInvoice(ctx, store.CreateInvoiceParams{
		Phone:     sql.NullString{String: phone, Valid: true},
		Reference: sql.NullString{String: chargeResponse.Data.Reference, Valid: true},
	}); err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	return chargeResponse, nil
}

func (p *PaystackServices) ReconcilePaystackMpesaCallback(payload PaystackCallbackResponse) error {
	invoiceStatus := model.InvoiceStatusProcessing
	data := payload.Data
	if payload.Event == "charge.success" && data.Status == "success" && data.Amount == 1500 {
		invoiceStatus = model.InvoiceStatusProcessed
		updatedInvoice, err := p.sqlStore.UpdateInvoiceForMpesa(ctx, store.UpdateInvoiceForMpesaParams{
			Reference:   sql.NullString{String: data.Reference, Valid: true},
			Msid:        sql.NullString{String: data.Authorization.MobileMoneyNumber, Valid: true},
			Channel:     sql.NullString{String: data.Authorization.Channel, Valid: true},
			Status:      invoiceStatus,
			Amount:      sql.NullString{String: strconv.Itoa(data.Amount), Valid: true},
			Bank:        sql.NullString{String: data.Authorization.Bank, Valid: true},
			AuthCode:    sql.NullString{String: data.Authorization.AuthCode, Valid: true},
			Fees:        sql.NullString{String: strconv.Itoa(data.Fees), Valid: true},
			CountryCode: sql.NullString{String: data.Authorization.CountryCode, Valid: true},
			Currency:    sql.NullString{String: data.Currency, Valid: true},
		})
		if err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return err
		}
		if _, err := p.sqlStore.UpdateLandlord(ctx, store.UpdateLandlordParams{
			IsLandlord: sql.NullBool{Bool: true, Valid: true},
			Phone:      updatedInvoice.Phone.String,
		}); err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return err
		}
	}
	return nil
}

func (p PaystackServices) ServiceName() string {
	return "PaystackServices"
}
