package entity

type TaskAssignee struct {
	ID     string
	TaskID string
	UserID string
}

type TaskAssignees []*TaskAssignee
