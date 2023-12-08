package services

import "github.com/sirupsen/logrus"

func NewEquityBankService(logger *logrus.Logger, env string) EquityService {
	return &equityClient{logger: logger, env: env}
}

type EquityService interface {
	ServiceName() string
}

type equityClient struct {
	logger *logrus.Logger
	env    string
}

func (*equityClient) ServiceName() string {
	return "equityClient"
}
