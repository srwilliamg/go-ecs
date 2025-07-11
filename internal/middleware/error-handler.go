package appMiddleware

import (
	"net/http"
	contextKey "srwilliamg/app/v1/internal/context-key"
	customError "srwilliamg/app/v1/internal/custom-error"
	"srwilliamg/app/v1/internal/request"

	"go.uber.org/zap"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(contextKey.LoggerKey).(*zap.SugaredLogger)
		defer func() {
			if err := recover(); err != nil {

				res, err := request.MarshalResponse[any](nil, customError.NewCustomError("Unexpected error", nil))

				if err != nil {
					logger.Error("Error marshalling response:", zap.Error(err))
					res = []byte(`{"error": "Internal Server Error"}`)
				}

				logger.Error("Unexpected error", zap.Any("error", err))
				request.PrepareResponse(&w, res, http.StatusInternalServerError)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
