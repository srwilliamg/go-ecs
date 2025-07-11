package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func UserRouter(app *chi.Mux) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of users"))
	})

	router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create a user"))
	})

	return router
}
