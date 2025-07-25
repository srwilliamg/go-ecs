package repository

import (
	"srwilliamg/app/v1/internal/domain/entities"
	dbInterface "srwilliamg/app/v1/internal/interfaces/db"
)

type Base struct {
	querier dbInterface.QuerierInterface[entities.User]
}

func (repo *Base) SetQuerier(db dbInterface.DatabaseInterface) {
	repo.querier = dbInterface.NewQuerier[entities.User](db)
}

func (repo *Base) GetQuerier() dbInterface.QuerierInterface[entities.User] {
	return repo.querier
}
