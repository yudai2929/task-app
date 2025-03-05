package entity

import "time"

type Task struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Status      TaskStatus
	DueDate     *time.Time
}

type TaskStatus int

const (
	TaskStatusUnknown TaskStatus = iota
	TaskStatusTodo
	TaskStatusInProgress
	TaskStatusDone
)

func (ts TaskStatus) String() string {
	switch ts {
	case TaskStatusUnknown:
		return "Unknown"
	case TaskStatusTodo:
		return "Todo"
	case TaskStatusInProgress:
		return "In Progress"
	case TaskStatusDone:
		return "Done"
	default:
		return "Unknown"
	}
}

func (ts TaskStatus) Int() int {
	return int(ts)
}

type Tasks []*Task

func (t *Task) Update(title, description string, status TaskStatus, dueDate *time.Time) {
	t.Title = title
	t.Description = description
	t.Status = status
	t.DueDate = dueDate
}
