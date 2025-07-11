package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func UserRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("List of users"))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Create a user"))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	return router
}
