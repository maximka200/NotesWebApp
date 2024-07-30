package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/maximka200/NotesWebApp"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(todo.User) (int, error) {
	return 0, nil
}
