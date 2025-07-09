package appMiddleware

import (
	"context"
	"net/http"
	contextKey "srwilliamg/app/v1/context-key"
	"srwilliamg/app/v1/logger"
)

func InitLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggerRequest := logger.GetLogger()
		ctx := context.WithValue(r.Context(), contextKey.LoggerKey, loggerRequest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
