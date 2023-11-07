package services

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/sirupsen/logrus"
)

type MpesaServices struct {
	transactionType     string
	logger              *logrus.Logger
	accessTokenEndpoint string
	lipaExpressEndpoint string
	callbackUrl         string
	client              *http.Client
	config              config.MpesaConfig
}

func NewMpesaService(callbackUrl string, transactionType string, cfg config.MpesaConfig, logger *logrus.Logger) *MpesaServices {
	return &MpesaServices{
		transactionType:     transactionType,
		logger:              logger,
		accessTokenEndpoint: fmt.Sprintf("%s/oauth/v1/generate?grant_type=client_credentials", cfg.BaseApi),
		lipaExpressEndpoint: fmt.Sprintf("%s/mpesa/stkpush/v1/processrequest", cfg.BaseApi),
		callbackUrl:         callbackUrl,
		client:              &http.Client{},
		config:              cfg,
	}
}

func (m *MpesaServices) GetAccessToken() string {
	dataToEncode := fmt.Sprintf("%s:%s", m.config.ConsumerKey, m.config.ConsumerSecret)
	sEnc := base64.StdEncoding.EncodeToString([]byte(dataToEncode))

	req, err := http.NewRequest("GET", m.accessTokenEndpoint, nil)
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", sEnc))

	res, err := m.client.Do(req)
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		m.logger.Errorf("%s:%v", m.ServiceName(), err)
	}
	fmt.Println(string(body))

	return "access_token"
}

func (m MpesaServices) ServiceName() string { return "MpesaServices" }
