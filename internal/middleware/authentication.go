package appMiddleware

import (
	"encoding/json"
	"net/http"
	contextKey "srwilliamg/app/v1/internal/context-key"
	customError "srwilliamg/app/v1/internal/custom-error"
	appRequest "srwilliamg/app/v1/internal/request"

	"go.uber.org/zap"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(contextKey.LoggerKey).(*zap.SugaredLogger)
		logger.Info("Request Auth:", zap.Any("request", r))

		var token string

		if r.Header.Get("x-api-key") != "" {
			token = r.Header.Get("x-api-key")
		}

		if r.Header.Get("Authorization") != "" {
			token = r.Header.Get("Authorization")
		}

		if token == "test" {
			next.ServeHTTP(w, r)
			return
		}

		res := appRequest.BaseResponse[any](nil, customError.NewCustomError("Authorization error", nil), nil)
		resJsonBytes, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		logger.Info("Response Auth:", zap.Any("response", res))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(resJsonBytes)
	})

}
