package repository

import "database/sql"

type authorization interface{}

type Repository struct {
	Auth authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth: newAuthRepository(db),
	}
}
