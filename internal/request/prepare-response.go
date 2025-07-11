package request

import "net/http"

func PrepareResponse(w *http.ResponseWriter, res []byte, code int) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(code)
	_, err := (*w).Write(res)

	if err != nil {
		http.Error(*w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
