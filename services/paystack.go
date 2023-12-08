package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/sirupsen/logrus"
)

type PaystackService interface {
	ChargeMpesaPhone(ctx context.Context, payload model.PaystackMpesaChargePayload) (*model.PaystackMpesaChargeResponse, error)
	ReconcilePaystackMpesaCallback(ctx context.Context, payload model.PaystackCallbackResponse) error
	ServiceName() string
}

type paystackClient struct {
	config   config.Paystack
	logger   *logrus.Logger
	baseApi  string
	env      string
	sqlStore *store.Queries
}

func NewPaystackService(cfg config.Paystack, env string, logger *logrus.Logger, sqlStore *store.Queries) PaystackService {
	return &paystackClient{config: cfg, logger: logger, env: env, baseApi: cfg.BaseApi, sqlStore: sqlStore}
}

func (p *paystackClient) ChargeMpesaPhone(ctx context.Context, payload model.PaystackMpesaChargePayload) (*model.PaystackMpesaChargeResponse, error) {
	phone := ctx.Value("phone").(string)
	var chargeResponse *model.PaystackMpesaChargeResponse

	url := p.baseApi + "/charge"
	payload.MobileMoney.Provider = "mpesa"
	if p.env == "development" {
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
		Msid:      sql.NullString{String: payload.MobileMoney.Phone, Valid: true},
		Phone:     sql.NullString{String: phone, Valid: true},
		Reference: sql.NullString{String: chargeResponse.Data.Reference, Valid: true},
	}); err != nil {
		p.logger.Errorf("%s:%v", p.ServiceName(), err)
		return nil, err
	}

	isLandlord := ctx.Value("is_landlord").(bool)
	if !isLandlord {
		if _, err := p.sqlStore.TrackSubscribeRetries(ctx, store.TrackSubscribeRetriesParams{
			Phone:            phone,
			SubscribeRetries: 1,
		}); err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return nil, err
		}
	}

	return chargeResponse, nil
}

func (p *paystackClient) ReconcilePaystackMpesaCallback(ctx context.Context, payload model.PaystackCallbackResponse) error {
	if payload.Event == "charge.success" {
		data := payload.Data

		createdAt, err := time.Parse(time.RFC3339, data.CreatedAt)
		if err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return err
		}

		paidAt, err := time.Parse(time.RFC3339, data.PaidAt)
		if err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return err
		}

		updatedInvoice, err := p.sqlStore.UpdateInvoiceForMpesa(ctx, store.UpdateInvoiceForMpesaParams{
			Reference:   sql.NullString{String: data.Reference, Valid: true},
			Channel:     sql.NullString{String: data.Authorization.Channel, Valid: true},
			Status:      model.InvoiceStatusProcessed.String(),
			Amount:      sql.NullString{String: strconv.Itoa(data.Amount / 100), Valid: true},
			Bank:        sql.NullString{String: data.Authorization.Bank, Valid: true},
			AuthCode:    sql.NullString{String: data.Authorization.AuthCode, Valid: true},
			Fees:        sql.NullString{String: strconv.Itoa(data.Fees / 100), Valid: true},
			CountryCode: sql.NullString{String: data.Authorization.CountryCode, Valid: true},
			Currency:    sql.NullString{String: data.Currency, Valid: true},
			CreatedAt:   createdAt,
			UpdatedAt:   paidAt,
		})
		if err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return err
		}

		if _, err := p.sqlStore.UpdateLandlord(ctx, store.UpdateLandlordParams{
			NextRenewal: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Phone:       updatedInvoice.Phone.String,
		}); err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return err
		}

		if _, err := p.sqlStore.TrackSubscribeRetries(ctx, store.TrackSubscribeRetriesParams{
			Phone:            updatedInvoice.Phone.String,
			SubscribeRetries: 0,
		}); err != nil {
			p.logger.Errorf("%s:%v", p.ServiceName(), err)
			return err
		}
	}
	return nil
}

func (p paystackClient) ServiceName() string {
	return "paystackClient"
}
