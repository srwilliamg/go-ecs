package main

import (
	"fmt"
	"net/http"
	"srwilliamg/app/v1/internal/application/routes"
	"srwilliamg/app/v1/internal/infrastructure/config"
	database "srwilliamg/app/v1/internal/infrastructure/db"
	"srwilliamg/app/v1/internal/infrastructure/logger"
	repository "srwilliamg/app/v1/internal/repositories/users"
	"time"

	appMiddleware "srwilliamg/app/v1/internal/application/middleware"

	"os"

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

	l := logger.GetLogger()
	config.Load()

	db, closer, err := database.Connect(l)
	if err != nil {
		os.Exit(1)
	}
	defer closer()

	repoUsers := &repository.UserRepository{}
	repoUsers.SetDB(db)

	app.Mount("/", routes.Routes(app))

	fmt.Fprintf(os.Stdout, "Starting server on %s\n", config.Envs.Port)

	err = http.ListenAndServe(":"+config.Envs.Port, app)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}

	fmt.Println("Server started successfully")
}

func main() {
	run()
}
