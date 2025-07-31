package appMiddleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	contextKey "srwilliamg/app/v1/internal/application/context-key"
	log "srwilliamg/app/v1/internal/interfaces/logger"
	"strings"
)

func RequestIdentifier(next http.Handler) http.Handler {
	handlerShowRequestIdentifier := func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(contextKey.LoggerKey).(*log.Logger)
		(*logger).Info("Request Identifier middleware invoked", log.String("method", r.Method))
		if r.Body != nil && strings.Compare(r.Method, "GET") != 0 {
			var buf bytes.Buffer
			var mapData map[string]any

			tee := io.TeeReader(r.Body, &buf)
			err := json.NewDecoder(tee).Decode(&mapData)

			if err != nil {
				(*logger).Error("TeeReader error", log.Any("error", err))
				next.ServeHTTP(w, r)
				return
			}

			r.Body = io.NopCloser(&buf)

			(*logger).Info("request content: ", log.Any("content", mapData))
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handlerShowRequestIdentifier)
}
