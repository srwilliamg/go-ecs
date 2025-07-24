package repositoryUsers

import (
	database "srwilliamg/app/v1/internal/interfaces/db"
	repository "srwilliamg/app/v1/internal/interfaces/repository"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	ID        int64   `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type UserRepository struct {
	repository.Base
}

func (repo UserRepository) GetUsers() (*database.Result, error) {
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
	results, err := repo.GetDB().Query(sql)

	return results, nil
}
