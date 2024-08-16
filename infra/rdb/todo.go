package rdb

import (
	"errors"

	"github.com/katsuokaisao/gin-play/domain"
	"github.com/katsuokaisao/gin-play/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	rdb *RDB
}

func NewTodoRepository(rdb *RDB) repository.TodoRepository {
	return &todoRepository{rdb: rdb}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	if err := r.rdb.NewSession(&gorm.Session{}).Create(todo).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) || errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrConflict
		}
	}
	return nil
}

func (r *todoRepository) List(todoFilter *domain.TodoFilter) ([]domain.Todo, error) {
	todos := make([]domain.Todo, 0)
	tx := r.rdb.NewSession(&gorm.Session{})
	if todoFilter.Assignee != nil {
		tx = tx.Where("assignee = ?", todoFilter.Assignee)
	}
	if todoFilter.Status != nil {
		tx = tx.Where("status = ?", todoFilter.Status)
	}
	if todoFilter.Priority != nil {
		tx = tx.Where("priority = ?", todoFilter.Priority)
	}
	if err := tx.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) Get(id int) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.rdb.NewSession(&gorm.Session{}).First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrRecordNotFound
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Update(id int, todo *domain.TodoUpdate) error {
	update := map[string]interface{}{}
	if todo.Title != nil {
		update["title"] = todo.Title
	}
	if todo.Assignee != nil {
		update["assignee"] = todo.Assignee
	}
	if todo.Status != nil {
		update["status"] = todo.Status
	}
	if todo.Priority != nil {
		update["priority"] = todo.Priority
	}
	if todo.BeginAt != nil {
		update["begin_at"] = todo.BeginAt
	}
	if todo.EndAt != nil {
		update["end_at"] = todo.EndAt
	}
	if todo.ExpireAt != nil {
		update["expire_at"] = todo.ExpireAt
	}
	if todo.Explanation != nil {
		update["explanation"] = todo.Explanation
	}

	result := r.rdb.NewSession(&gorm.Session{}).Model(&domain.Todo{}).Where("id = ?", id).Updates(update)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrRecordNotFound
	}
	return nil
}

func (r *todoRepository) Delete(id int) error {
	if err := r.rdb.NewSession(&gorm.Session{}).Delete(&domain.Todo{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrRecordNotFound
		}
		return err
	}
	return nil
}
