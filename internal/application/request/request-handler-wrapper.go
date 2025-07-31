package request

import (
	"net/http"
	"srwilliamg/app/v1/internal/interfaces/logger"
)

func WithReqHandlerWrapper(f func(w http.ResponseWriter, r *http.Request, logger *logger.Logger)) func(w http.ResponseWriter, r *http.Request) {

	wrappedRequestHandler := func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*logger.Logger)
		if !ok {
			(*logger).Warn("Logger not found in context")
		}

		f(w, r, logger)
	}

	return wrappedRequestHandler
}
