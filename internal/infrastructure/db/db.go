package database

import (
	"database/sql"

	"srwilliamg/app/v1/internal/infrastructure/config"
	database "srwilliamg/app/v1/internal/interfaces/db"
	l "srwilliamg/app/v1/internal/interfaces/logger"

	_ "github.com/lib/pq"
)

func Connect(logger *l.Logger) (database.DatabaseInterface, func(), error) {

	var connection = "postgres://" + config.Envs.DBUser + ":" + config.Envs.DBPass + "@" + config.Envs.DBHost + "/" + config.Envs.DBName + ":" + config.Envs.Port
	db, err := sql.Open("postgres", connection)

	if err != nil {
		(*logger).Error("Couldn't connect to database", l.Err(err))
		return nil, nil, err
	}
	closer := func() {
		db.Close()
	}

	return database.NewDatabase(db), closer, nil
}
