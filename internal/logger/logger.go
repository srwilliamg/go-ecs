package logger

import (
	"os"

	"go.uber.org/zap"
)

func GetLogger() *zap.SugaredLogger {
	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") != "production" {
		logger = zap.Must(zap.NewDevelopment())
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("Logger initialized")
	return sugar
}
