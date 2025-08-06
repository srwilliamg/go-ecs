package repositoryUsers

import (
	"srwilliamg/app/v1/internal/application/dto"
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

func (repo UserRepository) CreateUsers(insUsers []dto.User) (*database.Result[entities.User], error) {
	var err error
	sql := "insert into \"user\" (username, email, password, created_at, updated_at) values (:username, :email, :password, now(), now())"
	entitiesUsers := make([]*entities.User, len(insUsers))

	for i, v := range insUsers {
		v.Password, err = entities.HashPassword(v.Password)
		if err != nil {
			return nil, err
		}
		entitiesUsers[i] = v.ToEntity()
	}

	results, err := repo.GetQuerier().Mutate(sql, entitiesUsers)

	if err != nil {
		return nil, err
	}

	return results, nil
}
