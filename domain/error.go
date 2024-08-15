package domain

import "errors"

var (
	ErrTodoTitleRequired     = errors.New("todo title is required")
	ErrTodoTitleEmpty        = errors.New("todo title is empty")
	ErrTodoStatusInvalid     = errors.New("todo status is invalid")
	ErrTodoPriorityInvalid   = errors.New("todo priority is invalid")
	ErrTodoBeginAtAfterEndAt = errors.New("todo begin_at is after end_at")
	ErrTodoExpireAtBeforeNow = errors.New("todo expire_at is before now")
)
