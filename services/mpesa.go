package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/sirupsen/logrus"
)

func NewMpesaService(cfg config.MpesaConfig, logger *logrus.Logger, queries *store.Queries) MpesaService {
	return &mpesaClient{
		logger:              logger,
		accessTokenEndpoint: fmt.Sprintf("%s/oauth/v1/generate?grant_type=client_credentials", cfg.BaseApi),
		lipaExpressEndpoint: fmt.Sprintf("%s/mpesa/stkpush/v1/processrequest", cfg.BaseApi),
		config:              cfg,
		queries:             queries,
	}
}

type mpesaClient struct {
	logger              *logrus.Logger
	accessTokenEndpoint string
	lipaExpressEndpoint string
	config              config.MpesaConfig
	queries             *store.Queries
}

type MpesaService interface {
	GetAccessToken() (*model.AccessResponse, error)
	StkPush(payload model.LipaNaMpesaPayload) (*model.StkPushResponse, error)
	ServiceName() string
}

func (m *mpesaClient) GetAccessToken() (*model.AccessResponse, error) {
	response := &model.AccessResponse{}

	dataToEncode := fmt.Sprintf("%s:%s", m.config.ConsumerKey, m.config.ConsumerSecret)
	sEnc := base64.StdEncoding.EncodeToString([]byte(dataToEncode))

	req, err := http.NewRequest("GET", m.accessTokenEndpoint, nil)
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+sEnc)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}

	return response, nil
}

func (m *mpesaClient) StkPush(payload model.LipaNaMpesaPayload) (*model.StkPushResponse, error) {
	var stkResponse *model.StkPushResponse

	t := time.Now().Format("20060102150405")
	dataToEncode := fmt.Sprintf("%d%s%s", payload.BusinessShortCode, m.config.PassKey, t)
	password := base64.StdEncoding.EncodeToString([]byte(dataToEncode))
	payload.Timestamp = t
	payload.Password = password

	access, err := m.GetAccessToken()
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}
	req, err := http.NewRequest("POST", m.lipaExpressEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+access.AccessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&stkResponse); err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}

	/*resCode, err := strconv.Atoi(stkResponse.ResponseCode)
	if err != nil {
		m.logger.Errorf("%s:%v", "StkPushResponseCodeParsingError", err)
	}
	if stkResponse != nil && resCode == 0 {
		if _, err := m.queries.CreateInvoice(ctx, store.CreateInvoiceParams{
			WCoCheckoutID: sql.NullString{String: stkResponse.CheckoutRequestID, Valid: true},
			Reason:        sql.NullString{String: payload.TransactionDesc, Valid: true},
			Msid:          sql.NullString{String: strconv.Itoa(payload.PhoneNumber), Valid: true},
			Phone:         sql.NullString{String: strconv.Itoa(payload.PartyA), Valid: true},
		}); err != nil {
			return nil, err
		}
	}*/

	return stkResponse, nil
}

func (m mpesaClient) ServiceName() string { return "mpesaClient" }
