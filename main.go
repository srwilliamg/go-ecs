package main

import (
	"fmt"
	"net/http"
	"time"

	appMiddleware "srwilliamg/app/v1/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func createApp() *chi.Mux {
	app := chi.NewRouter()
	app.Use(middleware.Timeout(60 * time.Second))
	app.Use(appMiddleware.InitLogger, appMiddleware.ErrorHandler, appMiddleware.Auth, appMiddleware.RequestIdentifier)
	return app
}

func run() {
	app := createApp()

	app.Get("/", func(w http.ResponseWriter, app *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", app)
}

func main() {
	run()
}
