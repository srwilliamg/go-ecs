package appMiddleware

import (
	"net/http"
	contextKey "srwilliamg/app/v1/internal/application/context-key"
	customError "srwilliamg/app/v1/internal/application/custom-error"
	"srwilliamg/app/v1/internal/application/request"
	"srwilliamg/app/v1/internal/interfaces/logger"
	log "srwilliamg/app/v1/internal/interfaces/logger"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(contextKey.LoggerKey).(*logger.Logger)
		(*logger).Info("Request Auth:", log.Any("request", r))

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

		res, err := request.MarshalResponse[any](nil, customError.NewCustomError("You are not Authorized", nil))
		if err != nil {
			(*logger).Error("Error marshalling response:", log.Err(err))
		}

		(*logger).Info("Response Auth:", log.Any("response", res))
		request.PrepareResponse(&w, res, http.StatusUnauthorized)
	})

}
