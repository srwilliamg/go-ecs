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
	defer (func() {
		err := logger.Sync()
		if err != nil {
			_, err := os.Stderr.WriteString("Failed to sync logger: " + err.Error() + "\n")
			if err != nil {
				_, _ = os.Stderr.WriteString("Error writing to stderr: " + err.Error() + "\n")
				os.Exit(1)
			} else {
				_, _ = os.Stderr.WriteString("Logger synced successfully\n")
			}
		}
	})()
	sugar := logger.Sugar()

	sugar.Info("Logger initialized")
	return sugar
}
