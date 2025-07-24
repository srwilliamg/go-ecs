package repository

import dbInterface "srwilliamg/app/v1/internal/interfaces/db"

type Base struct {
	db dbInterface.DatabaseInterface
}

func (repo *Base) SetDB(db dbInterface.DatabaseInterface) { repo.db = db }

func (repo *Base) GetDB() dbInterface.DatabaseInterface {
	if repo.db == nil {
		return nil
	}
	return repo.db
}
