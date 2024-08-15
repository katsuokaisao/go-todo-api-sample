package usecase

import (
	"github.com/katsuokaisao/gin-play/domain"
	"github.com/katsuokaisao/gin-play/repository"
)

type TodoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUseCase(todoRepository repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		todoRepository: todoRepository,
	}
}

func (u *TodoUseCase) Create(todo *domain.Todo) error {
	return u.todoRepository.Create(todo)
}

func (u *TodoUseCase) List(todoFilter *domain.TodoFilter) ([]domain.Todo, error) {
	return u.todoRepository.List(todoFilter)
}

func (u *TodoUseCase) Get(id int) (*domain.Todo, error) {
	return u.todoRepository.Get(id)
}

func (u *TodoUseCase) Update(id int, todo *domain.TodoUpdate) error {
	return u.todoRepository.Update(id, todo)
}

func (u *TodoUseCase) Delete(id int) error {
	return u.todoRepository.Delete(id)
}
