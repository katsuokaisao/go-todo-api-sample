package domain

import "time"

type TodoStatus string

const (
	TodoStatusOpen       TodoStatus = "open"
	TodoStatusProcessing TodoStatus = "processing"
	TodoStatusPending    TodoStatus = "pending"
	TodoStatusDone       TodoStatus = "done"
	TodoStatusUnknown    TodoStatus = ""
)

func TodoStatusFromString(s string) TodoStatus {
	switch s {
	case "open":
		return TodoStatusOpen
	case "processing":
		return TodoStatusProcessing
	case "pending":
		return TodoStatusPending
	case "done":
		return TodoStatusDone
	default:
		return TodoStatusUnknown
	}
}

type TodoPriority string

const (
	TodoPriorityLow     TodoPriority = "low"
	TodoPriorityNormal  TodoPriority = "normal"
	TodoPriorityHigh    TodoPriority = "high"
	TodoPriorityUnknown TodoPriority = ""
)

func TodoPriorityFromString(s string) TodoPriority {
	switch s {
	case "low":
		return TodoPriorityLow
	case "normal":
		return TodoPriorityNormal
	case "high":
		return TodoPriorityHigh
	default:
		return TodoPriorityUnknown
	}
}

type Todo struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Assignee    *string       `json:"assignee"`
	Status      TodoStatus    `json:"status"`
	Priority    *TodoPriority `json:"priority"`
	BeginAt     *time.Time    `json:"begin_at"`
	EndAt       *time.Time    `json:"end_at"`
	ExpireAt    *time.Time    `json:"expire_at"`
	Expired     *bool         `json:"expired"`
	Explanation *string       `json:"explanation"`
}

type TodoFilter struct {
	Assignee *string
	Status   *TodoStatus
	Priority *TodoPriority
}

type TodoUpdate struct {
	Title       *string
	Assignee    *string
	Status      *TodoStatus
	Priority    *TodoPriority
	BeginAt     *time.Time
	EndAt       *time.Time
	ExpireAt    *time.Time
	Explanation *string
}
