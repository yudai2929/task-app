package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/google/uuid"
	xo "github.com/yudai2929/task-app/database/gen"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
)

type taskAssigneeRepository struct {
	db *sql.DB
}

func NewTaskAssigneeRepository(db *sql.DB) *taskAssigneeRepository {
	return &taskAssigneeRepository{
		db: db,
	}
}

func (r *taskAssigneeRepository) batchCreate(ctx context.Context, tx *sql.Tx, taskID string, userIDs []string) error {
	if len(userIDs) == 0 {
		return nil
	}

	assignees := make([]*xo.TaskAssignee, len(userIDs))
	for i, userID := range userIDs {
		assignees[i] = &xo.TaskAssignee{
			ID:     uuid.NewString(),
			TaskID: taskID,
			UserID: userID,
		}
	}

	values := make([]string, 0, len(assignees))
	args := make([]interface{}, 0, len(assignees)*5)
	for i, assignee := range assignees {
		values = append(values, `($`+strconv.Itoa(i*5+1)+`, $`+strconv.Itoa(i*5+2)+`, $`+strconv.Itoa(i*5+3)+`, $`+strconv.Itoa(i*5+4)+`, $`+strconv.Itoa(i*5+5)+`)`)
		args = append(args, assignee.ID, assignee.TaskID, assignee.UserID, assignee.CreatedAt, assignee.UpdatedAt)
	}

	query := `
		INSERT INTO task_assignees (id, task_id, user_id, created_at, updated_at)
		VALUES ` + strings.Join(values, ",")
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Convert(err)
	}
	return nil
}

func (r *taskAssigneeRepository) batchDeleteByTaskID(ctx context.Context, tx *sql.Tx, taskID string) error {
	_, err := tx.ExecContext(ctx, `
	DELETE FROM task_assignees
	WHERE task_id = $1`,
		taskID)
	if err != nil {
		return errors.Convert(err)
	}
	return nil
}

func (r *taskAssigneeRepository) UpdateTaskAssignees(ctx context.Context, taskID string, userIDs []string) error {
	return runInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		if err := r.batchDeleteByTaskID(ctx, tx, taskID); err != nil {
			return err
		}
		if err := r.batchCreate(ctx, tx, taskID, userIDs); err != nil {
			return err
		}
		return nil
	})
}

func (r *taskAssigneeRepository) GetTaskAssignee(ctx context.Context, taskID, userID string) (*entity.TaskAssignee, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, task_id, user_id
		FROM task_assignees
		WHERE task_id = $1 AND user_id = $2`,
		taskID, userID)

	assignee := &entity.TaskAssignee{}
	if err := row.Scan(&assignee.ID, &assignee.TaskID, &assignee.UserID); err != nil {
		return nil, errors.Convert(err)
	}
	return assignee, nil
}
