package postgres

import (
	"database/sql"
)

type RepositoryHandler interface {
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryHandler {
	return &repository{
		db: db,
	}
}
