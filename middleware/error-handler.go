package appMiddleware

import (
	"encoding/json"
	"net/http"
	contextKey "srwilliamg/app/v1/context-key"
	customError "srwilliamg/app/v1/custom-error"
	appRequest "srwilliamg/app/v1/request"

	"go.uber.org/zap"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(contextKey.LoggerKey).(*zap.SugaredLogger)
		defer func() {
			if err := recover(); err != nil {
				res := appRequest.BaseResponse[any](nil, customError.NewCustomError("Unexpected error", nil), nil)
				resJsonBytes, marshalErr := json.Marshal(res)
				if marshalErr != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				logger.Error("Unexpected error", zap.Any("error", err))
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resJsonBytes)

				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
