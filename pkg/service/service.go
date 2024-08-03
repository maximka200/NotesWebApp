package service

import (
	todo "github.com/maximka200/NotesWebApp"
	"github.com/maximka200/NotesWebApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	CreateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
