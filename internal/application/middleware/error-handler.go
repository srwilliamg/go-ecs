package appMiddleware

import (
	"net/http"
	contextKey "srwilliamg/app/v1/internal/application/context-key"
	customError "srwilliamg/app/v1/internal/application/custom-error"
	"srwilliamg/app/v1/internal/application/request"
	"srwilliamg/app/v1/internal/interfaces/logger"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := r.Context().Value(contextKey.LoggerKey).(logger.Logger)
		defer func() {
			if err := recover(); err != nil {

				(log).Error("Error Detected", logger.Any("Error", err))

				log.Error("Unexpected error", logger.Any("Error", err))
				request.PrepareResponse(&w, customError.NewCustomError("Unexpected error", nil), http.StatusInternalServerError, log)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
