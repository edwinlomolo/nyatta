package services

import (
	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"go.uber.org/zap"
)

func NewLogger(cfg *nyatta_context.Config) (log *zap.SugaredLogger, err error) {
	var logger *zap.SugaredLogger
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
