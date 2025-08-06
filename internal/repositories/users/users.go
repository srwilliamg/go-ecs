package repositoryUsers

import (
	"srwilliamg/app/v1/internal/domain/entities"
	database "srwilliamg/app/v1/internal/interfaces/db"
	repository "srwilliamg/app/v1/internal/interfaces/repository"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	repository.Base
}

func (repo UserRepository) GetUsers() (*database.Result[entities.User], error) {
	getUsersQuery := func() (string, error) {
		users := sq.Select("*").From("\"user\" u")
		activeUsers := users //.Where(sq.Eq{"deleted_at": nil})

		sql, _, err := activeUsers.ToSql()
		if err != nil {
			return "", err
		}

		return sql, nil
	}

	sql, err := getUsersQuery()

	if err != nil {
		return nil, err
	}

	scanUser := func() entities.User {
		return entities.User{}
	}

	results, err := repo.GetQuerier().Query(sql, scanUser)

	if err != nil {
		return nil, err
	}

	return results, nil
}
