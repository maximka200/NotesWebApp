package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/maximka200/NotesWebApp"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
