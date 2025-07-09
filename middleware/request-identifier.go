package appMiddleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	contextKey "srwilliamg/app/v1/context-key"
	"strings"

	"go.uber.org/zap"
)

func RequestIdentifier(next http.Handler) http.Handler {
	handlerShowRequestIdentifier := func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(contextKey.LoggerKey).(*zap.SugaredLogger)
		logger.Info("Request Identifier middleware invoked", zap.String("method", r.Method))
		if r.Body != nil && strings.Compare(r.Method, "GET") != 0 {
			var buf bytes.Buffer
			var mapData map[string]interface{}

			tee := io.TeeReader(r.Body, &buf)
			err := json.NewDecoder(tee).Decode(&mapData)

			if err != nil {
				logger.Error("TeeReader error", zap.Any("error", err))
				next.ServeHTTP(w, r)
				return
			}

			r.Body = io.NopCloser(&buf)

			logger.Info("request content: ", zap.Any("content", mapData))
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handlerShowRequestIdentifier)
}
