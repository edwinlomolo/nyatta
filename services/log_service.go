package service

import (
	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"go.uber.org/zap"
)

func NewLogger(cfg *nyatta_context.Config) (log *zap.Logger, err error) {
	var logger *zap.Logger
	if cfg.Env == "development" || cfg.Env == "test" {
		newLogger, err := zap.NewDevelopment()
		if err = err; err != nil {
			return nil, err
		}
		logger = newLogger
		return logger, nil
	}
	newLogger, err := zap.NewProduction()
	if err = err; err != nil {
		return nil, err
	}
	logger = newLogger
	return logger, nil
}
