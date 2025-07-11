package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(app *chi.Mux) {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	router.Mount("/v1/users", UserRouter(app))
}
