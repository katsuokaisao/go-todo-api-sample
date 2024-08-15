package repository

import "github.com/katsuokaisao/gin-play/domain"

type TodoRepository interface {
	Create(todo *domain.Todo) error
	List(todoFilter *domain.TodoFilter) ([]domain.Todo, error)
	Get(id int) (*domain.Todo, error)
	Update(id int, todo *domain.TodoUpdate) error
	Delete(id int) error
}
