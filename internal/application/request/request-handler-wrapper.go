package request

import (
	"net/http"
	contextKey "srwilliamg/app/v1/internal/application/context-key"
	log "srwilliamg/app/v1/internal/interfaces/logger"
)

func WithReqHandlerWrapper(f func(w http.ResponseWriter, r *http.Request, logger log.Logger)) func(w http.ResponseWriter, r *http.Request) {

	wrappedRequestHandler := func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value(contextKey.LoggerKey).(log.Logger)

		if !ok {
			if logger != nil {
				(logger).Warn("Logger not found in context")
			}
		}

		f(w, r, logger)
	}

	return wrappedRequestHandler
}
