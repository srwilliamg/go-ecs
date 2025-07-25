package database

import (
	"database/sql"
)

type DatabaseInterface interface {
	SetDB(db *sql.DB)
	Query(rawSql string) (*Result[any], error)
	GetDB() *sql.DB
}

type Database struct {
	db *sql.DB
}

func (rdb *Database) SetDB(db *sql.DB) {
	rdb.db = db
}
func (rdb *Database) GetDB() *sql.DB {
	return rdb.db
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

type Result[T any] struct {
	Rows []T
}

type QuerierInterface[T any] interface {
	GetDB() *sql.DB
	Query(rawSql string, scanFunc func([]any) T) (*Result[T], error)
}

type Querier[T any] struct {
	DatabaseInterface
}

func NewQuerier[T any](db DatabaseInterface) *Querier[T] {
	return &Querier[T]{
		DatabaseInterface: db,
	}
}

func (q *Querier[T]) GetDB() *sql.DB {
	return q.DatabaseInterface.GetDB()
}

func (querier *Querier[T]) Query(rawSql string, scanFunc func([]any) T) (*Result[T], error) {
	rows, err := querier.GetDB().Query(rawSql)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := &Result[T]{}
	for rows.Next() {
		values := make([]any, len(cols))
		pointers := make([]any, len(cols))
		for i := range values {
			pointers[i] = &values[i]
		}
		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}
		result.Rows = append(result.Rows, scanFunc(values))
	}
	return result, nil
}
