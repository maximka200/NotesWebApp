package service

import 
repository "github.com/maximka200/NotesWebApp/pkg/repository"

type Authorization interface {

}

type TodoList interface {
	
}

type TodoItem interface {
	
}

type Servis struct {
	Authorization
	TodoItem
	TodoList
}

type NewService(repos *repository.Repository) *Service {
	return &Service{}
}