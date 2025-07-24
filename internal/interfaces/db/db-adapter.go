package database

import "database/sql"

type DatabaseInterface interface {
	SetDB(db *sql.DB)
	Query(rawSql string) (*Result, error)
}

type Database struct {
	db *sql.DB
}

func (rdb *Database) SetDB(db *sql.DB) {
	rdb.db = db
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

type Result struct {
	rows []any
}

func (db *Database) Query(rawSql string) (*Result, error) {
	rows, err := db.db.Query(rawSql)
	if err != nil {
		return nil, err
	}

	result := &Result{
		rows: make([]any, 0),
	}

	for rows.Next() {
		var row any
		if err := rows.Scan(&row); err != nil {
			return nil, err
		}
		result.rows = append(result.rows, row)
	}
	return result, nil
}
