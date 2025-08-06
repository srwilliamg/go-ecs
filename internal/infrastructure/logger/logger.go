package logger

import (
	"os"
	"srwilliamg/app/v1/internal/interfaces/logger"

	"go.uber.org/zap"
)

func GetLogger() logger.Logger {
	zapLog := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") != "production" {
		zapLog = zap.Must(zap.NewDevelopment())
	}
	zapLogger := logger.NewZapAdapter(zapLog)
	defer (func() {
		err := zapLogger.Sync()
		if err != nil {
			_, err := os.Stderr.WriteString("Failed to sync logger: " + err.Error() + "\n")
			if err != nil {
				_, _ = os.Stderr.WriteString("Error writing to stderr: " + err.Error() + "\n")
				os.Exit(1)
			} else {
				_, _ = os.Stderr.WriteString("zapLogger synced successfully\n")
			}
		}
	})()

	zapLogger.Info("Logger initialized")
	return zapLogger
}
