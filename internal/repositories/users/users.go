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

		users := sq.Select("*").From("users")
		activeUsers := users.Where(sq.Eq{"deleted_at": nil})

		sql, _, err := activeUsers.ToSql()
		if err != nil {
			return "", err
		}

		return sql, nil
	}

	sql, err := getUsersQuery()

	if err == nil {
		return nil, err
	}

	scanUser := func(values []any) entities.User {
		return entities.User{
			ID:        values[0].(int64),   // ID
			Username:  values[1].(string),  // Username
			Email:     values[2].(string),  // Email
			Password:  values[3].(string),  // Password
			CreatedAt: values[4].(string),  // CreatedAt
			UpdatedAt: values[5].(string),  // UpdatedAt
			DeletedAt: values[6].(*string), // DeletedAt (can be nil)

		}
	}

	results, err := repo.GetQuerier().Query(sql, scanUser)

	return results, nil
}
