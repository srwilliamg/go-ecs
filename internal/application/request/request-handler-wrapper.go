package request

import (
	"net/http"

	zap "go.uber.org/zap"
)

func WithReqHandlerWrapper(f func(w http.ResponseWriter, r *http.Request, logger *zap.Logger)) func(w http.ResponseWriter, r *http.Request) {

	wrappedRequestHandler := func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*zap.Logger)
		if !ok {
			logger.Warn("Logger not found in context")
		}

		f(w, r, logger)
	}

	return wrappedRequestHandler
}
