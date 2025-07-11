package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(app *chi.Mux) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	router.Mount("/v1/users", UserRouter())
	return router
}
