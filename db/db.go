package database

import (
	"database/sql"

	"srwilliamg/app/v1/internal/config"
	l "srwilliamg/app/v1/internal/logger"
)

func Connect(logger *l.Logger) (*sql.DB, func(), error) {

	var connection = "postgres://" + config.Envs.DBUser + ":" + config.Envs.DBPass + "@" + config.Envs.DBHost + "/" + config.Envs.DBName + ":" + config.Envs.Port
	db, err := sql.Open("postgres", connection)

	if err != nil {
		(*logger).Error("Couldn't connect to database", l.Err(err))
		return nil, nil, err
	}
	closer := func() {
		db.Close()
	}

	return db, closer, nil
}
