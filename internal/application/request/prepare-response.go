package request

import (
	"net/http"
	"srwilliamg/app/v1/internal/interfaces/logger"
)

func PrepareResponse[T any](w *http.ResponseWriter, body T, code int, log logger.Logger) {
	res, err := MarshalResponse(body)

	if err != nil {
		log.Error("Error marshalling response:", logger.Err(err))
		res = []byte(`{"error": "Internal Server Error"}`)
	}

	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(code)
	_, err = (*w).Write(res)

	if err != nil {
		http.Error(*w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
