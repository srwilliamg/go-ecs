package main

import (
	"fmt"
	"net/http"
	"time"

	"srwilliamg/app/v1/internal/config"
	appMiddleware "srwilliamg/app/v1/internal/middleware"
	"srwilliamg/app/v1/routes"

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
	config.Load()
	app.Mount("/", routes.Routes(app))

	fmt.Printf("Starting server on %s\n", config.Envs.Port)
	err := http.ListenAndServe(":"+config.Envs.Port, app)

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	fmt.Println("Server started successfully")
}

func main() {
	run()
}
