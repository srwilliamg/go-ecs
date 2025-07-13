package repositories

import (
	sq "github.com/Masterminds/squirrel"
)

func GetUsersQuery() (string, error) {
	users := sq.Select("*").From("users")
	activeUsers := users.Where(sq.Eq{"deleted_at": nil})

	sql, _, err := activeUsers.ToSql()
	if err != nil {
		return "", err
	}

	return sql, nil
}
