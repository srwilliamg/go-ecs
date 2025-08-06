package appMiddleware

import (
	"fmt"
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

				fmt.Println("ERORRORORORORO:    ", fmt.Errorf("errr", err))

				res, err := request.MarshalResponse[any](nil, customError.NewCustomError("Unexpected error", nil))

				if err != nil {
					log.Error("Error marshalling response:", logger.Err(err))
					res = []byte(`{"error": "Internal Server Error"}`)
				}

				log.Error("Unexpected error", logger.Err(err))
				request.PrepareResponse(&w, res, http.StatusInternalServerError)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
