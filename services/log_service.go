package services

import (
	"github.com/3dw1nM0535/nyatta/config"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func NewLogger(cfg *config.Config) (log *zap.SugaredLogger, err error) {
	if cfg.Env == "development" || cfg.Env == "test" {
		newLogger, err := zap.NewDevelopment()
		if err = err; err != nil {
			return nil, err
		}
		logger = newLogger.Sugar()
	} else {
		newLogger, err := zap.NewProduction()
		if err = err; err != nil {
			return nil, err
		}

		logger = newLogger.Sugar()
	}

	defer logger.Sync()
	return logger, nil
}
