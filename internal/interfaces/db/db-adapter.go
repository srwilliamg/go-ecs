package database

import (
	"github.com/jmoiron/sqlx"
)

type DatabaseInterface interface {
	SetDB(db *sqlx.DB)
	GetDB() *sqlx.DB
}

type Database struct {
	db *sqlx.DB
}

func (rdb *Database) SetDB(db *sqlx.DB) {
	rdb.db = db
}
func (rdb *Database) GetDB() *sqlx.DB {
	return rdb.db
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{db: db}
}

type Result[T any] struct {
	Rows []T
}

type QuerierInterface[T any] interface {
	GetDB() *sqlx.DB
	Query(rawSql string, scanFunc func() T) (*Result[T], error)
}

type Querier[T any] struct {
	DatabaseInterface
}

func NewQuerier[T any](db DatabaseInterface) *Querier[T] {
	return &Querier[T]{
		DatabaseInterface: db,
	}
}

func (q *Querier[T]) GetDB() *sqlx.DB {
	return q.DatabaseInterface.GetDB()
}

func (querier *Querier[T]) Query(rawSql string, scanFunc func() T) (*Result[T], error) {
	rows, err := querier.GetDB().Queryx(rawSql)
	if err != nil {
		return nil, err
	}

	result := &Result[T]{}
	for rows.Next() {
		val := scanFunc()
		err := rows.StructScan(&val)
		if err != nil {
			return nil, err
		}
		result.Rows = append(result.Rows, val)
	}

	return result, nil
}
