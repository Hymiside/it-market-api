package repository

import "database/sql"

type authRepository struct {
	db *sql.DB
}

func newAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db: db}
}
