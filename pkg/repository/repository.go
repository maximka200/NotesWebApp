package repository



type Authorization interface {

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

type NewRepository() *Repository {
	return &Repository{}
}