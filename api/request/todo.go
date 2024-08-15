package request

import (
	"time"

	"github.com/katsuokaisao/gin-play/domain"
)

type TodoCreateRequest struct {
	Title       string     `json:"title"`
	Assignee    *string    `json:"assignee"`
	Status      *string    `json:"status"`
	Priority    *string    `json:"priority"`
	BeginAt     *time.Time `json:"begin_at"`
	EndAt       *time.Time `json:"end_at"`
	ExpireAt    *time.Time `json:"expire_at"`
	Explanation *string    `json:"explanation"`
}

func (r *TodoCreateRequest) Validate() error {
	if r.Title == "" {
		return domain.ErrTodoTitleRequired
	}
	if r.Status != nil {
		ts := domain.TodoStatus(*r.Status)
		if ts == domain.TodoStatusUnknown {
			return domain.ErrTodoStatusInvalid
		}
	}
	if r.Priority != nil {
		tp := domain.TodoPriorityFromString(*r.Priority)
		if tp == domain.TodoPriorityUnknown {
			return domain.ErrTodoPriorityInvalid
		}
	}
	if r.BeginAt != nil && r.EndAt != nil {
		if r.BeginAt.After(*r.EndAt) {
			return domain.ErrTodoBeginAtAfterEndAt
		}
	}
	if r.ExpireAt != nil {
		if r.ExpireAt.Before(time.Now()) {
			return domain.ErrTodoExpireAtBeforeNow
		}
	}

	return nil
}

func (r *TodoCreateRequest) ToDomain() *domain.Todo {
	d := &domain.Todo{
		Title:       r.Title,
		Assignee:    r.Assignee,
		BeginAt:     r.BeginAt,
		EndAt:       r.EndAt,
		ExpireAt:    r.ExpireAt,
		Explanation: r.Explanation,
	}
	if r.Status != nil {
		d.Status = domain.TodoStatus(*r.Status)
		if d.Status != domain.TodoStatusDone {
			expired := false
			if r.ExpireAt != nil && r.ExpireAt.Before(time.Now()) {
				expired = true
			}
			d.Expired = &expired
		}
	} else {
		d.Status = domain.TodoStatusOpen
	}
	if r.Priority != nil {
		p := domain.TodoPriorityFromString(*r.Priority)
		d.Priority = &p
	}
	return d
}

type TodoUpdateRequest struct {
	Title       *string    `json:"title"`
	Assignee    *string    `json:"assignee"`
	Status      *string    `json:"status"`
	Priority    *string    `json:"priority"`
	BeginAt     *time.Time `json:"begin_at"`
	EndAt       *time.Time `json:"end_at"`
	ExpireAt    *time.Time `json:"expire_at"`
	Explanation *string    `json:"explanation"`
}

func (r *TodoUpdateRequest) Validate() error {
	if r.Title != nil && *r.Title == "" {
		return domain.ErrTodoTitleRequired
	}
	if r.Status != nil {
		ts := domain.TodoStatus(*r.Status)
		if ts == domain.TodoStatusUnknown {
			return domain.ErrTodoStatusInvalid
		}
	}
	if r.Priority != nil {
		tp := domain.TodoPriorityFromString(*r.Priority)
		if tp == domain.TodoPriorityUnknown {
			return domain.ErrTodoPriorityInvalid
		}
	}
	if r.BeginAt != nil && r.EndAt != nil {
		if r.BeginAt.After(*r.EndAt) {
			return domain.ErrTodoBeginAtAfterEndAt
		}
	}
	if r.ExpireAt != nil {
		if r.ExpireAt.Before(time.Now()) {
			return domain.ErrTodoExpireAtBeforeNow
		}
	}
	return nil
}

func (r *TodoUpdateRequest) ToDomain() *domain.TodoUpdate {
	d := &domain.TodoUpdate{
		Title:       r.Title,
		Assignee:    r.Assignee,
		BeginAt:     r.BeginAt,
		EndAt:       r.EndAt,
		ExpireAt:    r.ExpireAt,
		Explanation: r.Explanation,
	}
	if r.Status != nil {
		ts := domain.TodoStatus(*r.Status)
		d.Status = &ts
	}
	if r.Priority != nil {
		tp := domain.TodoPriorityFromString(*r.Priority)
		d.Priority = &tp
	}
	return d
}

type TodoListRequest struct {
	Assignee *string `form:"assignee"`
	Status   *string `form:"status"`
	Priority *string `form:"priority"`
}

func (r *TodoListRequest) Validate() error {
	if r.Status != nil {
		ts := domain.TodoStatus(*r.Status)
		if ts == domain.TodoStatusUnknown {
			return domain.ErrTodoStatusInvalid
		}
	}
	if r.Priority != nil {
		tp := domain.TodoPriorityFromString(*r.Priority)
		if tp == domain.TodoPriorityUnknown {
			return domain.ErrTodoPriorityInvalid
		}
	}
	return nil
}

func (r *TodoListRequest) ToDomain() *domain.TodoFilter {
	f := &domain.TodoFilter{
		Assignee: r.Assignee,
	}
	if r.Status != nil {
		ts := domain.TodoStatus(*r.Status)
		f.Status = &ts
	}
	if r.Priority != nil {
		tp := domain.TodoPriorityFromString(*r.Priority)
		f.Priority = &tp
	}
	return f
}

type TodoGetRequest struct {
	ID int `uri:"id" binding:"required"`
}

type TodoDeleteRequest struct {
	ID int `uri:"id" binding:"required"`
}
