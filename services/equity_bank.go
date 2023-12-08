package services

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
	"net/http"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/sirupsen/logrus"
)

func NewEquityBankService(logger *logrus.Logger, env string, config config.EquityBank) EquityService {
	return &equityClient{logger: logger, config: config, env: env}
}

type EquityService interface {
	ServiceName() string
	AccountBalance(accountId, countryCode string) (*model.AccountBalanceResponse, error)
}

type equityClient struct {
	logger *logrus.Logger
	config config.EquityBank
	env    string
}

func (eq *equityClient) decodePrivateKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(eq.config.PrivateKey))
	if block == nil || block.Type != "PRIVATE KEY" {
		reason := "failed to decode PEM block containing private key"
		log.Fatal(reason)
		return nil, errors.New(reason)
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Println("Error decoding private key")
		return nil, err
	}

	return (key).(*rsa.PrivateKey), nil
}

func (eq *equityClient) authenticate() (*model.AuthResponse, error) {
	var authRes *model.AuthResponse

	api := eq.config.BaseApi + "/authentication/api/v3/authenticate/merchant"
	payload, err := json.Marshal(&model.AuthRequest{
		ConsumerSecret: eq.config.ConsumerSecret,
		MerchantCode:   eq.config.MerchantCode,
	})
	if err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(payload))
	if err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Api-Key", eq.config.ApiKey)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&authRes); err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}

	return authRes, nil
}

func (eq *equityClient) signMessageToBase64Encode(payload string) (string, error) {
	message := []byte(payload)
	hashedMessage := sha256.Sum256(message)
	privateKey, err := eq.decodePrivateKey()
	if err != nil {
		return "", err
	}

	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashedMessage[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func (eq *equityClient) AccountBalance(accountId, countryCode string) (*model.AccountBalanceResponse, error) {
	var accBalance *model.AccountBalanceResponse
	api := eq.config.BaseApi + "/v3-apis/account-api/v3.0/accounts/balances/" + countryCode + "/" + accountId
	reqSignature, err := eq.signMessageToBase64Encode(countryCode + accountId)
	if err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}
	authToken, err := eq.authenticate()
	if err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}
	req.Header.Add("Authorization", authToken.AccessToken)
	req.Header.Add("signature", reqSignature)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&accBalance); err != nil {
		eq.logger.Errorf("%s:%v", eq.ServiceName(), err)
		return nil, err
	}

	return accBalance, nil
}

func (*equityClient) ServiceName() string {
	return "equityClient"
}
