package database

import (
	"srwilliamg/app/v1/internal/infrastructure/config"
	database "srwilliamg/app/v1/internal/interfaces/db"
	l "srwilliamg/app/v1/internal/interfaces/logger"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func Connect(logger l.Logger) (database.DatabaseInterface, func(), error) {

	var connection = "postgres://" + config.Envs.DBUser + ":" + config.Envs.DBPass + "@" + config.Envs.DBHost + ":" + config.Envs.DBPort + "/" + config.Envs.DBName + "?sslmode=disable"
	logger.Info(connection)
	db, err := sqlx.Connect("postgres", connection)

	if err != nil {
		(logger).Error("Couldn't connect to database", l.Err(err))
		return nil, nil, err
	}
	closer := func() {
		db.Close()
	}

	return database.NewDatabase(db), closer, nil
}
