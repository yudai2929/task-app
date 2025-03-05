package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	xo "github.com/yudai2929/task-app/database/gen"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
)

type taskAssigneeRepository struct {
	db *sql.DB
}

func NewTaskAssigneeRepository(db *sql.DB) *taskAssigneeRepository {
	return &taskAssigneeRepository{
		db: db,
	}
}

func (r *taskAssigneeRepository) BatchCreate(ctx context.Context, assignees entity.TaskAssignees) error {
	db := getDB(ctx, r.db)
	now := time.Now()

	var values []string
	var args []interface{}
	for i, assignee := range assignees {
		values = append(values, `($`+strconv.Itoa(i*5+1)+`, $`+strconv.Itoa(i*5+2)+`, $`+strconv.Itoa(i*5+3)+`, $`+strconv.Itoa(i*5+4)+`, $`+strconv.Itoa(i*5+5)+`)`)
		args = append(args, assignee.ID, assignee.TaskID, assignee.UserID, now, now)
	}

	query := `INSERT INTO task_assignees (id, task_id, user_id, created_at, updated_at) VALUES ` + strings.Join(values, ",")
	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Convert(err)
	}
	return nil
}

func (r *taskAssigneeRepository) BatchDeleteByTaskID(ctx context.Context, taskID string) error {
	db := getDB(ctx, r.db)
	const sqlstr = `DELETE FROM task_assignees WHERE task_id = $1`
	_, err := db.ExecContext(ctx, sqlstr, taskID)
	if err != nil {
		return errors.Convert(err)
	}
	return nil
}

func (r *taskAssigneeRepository) GetTaskAssignee(ctx context.Context, taskID, userID string) (*entity.TaskAssignee, error) {
	db := getDB(ctx, r.db)
	const sqlstr = `SELECT id, task_id, user_id, created_at, updated_at FROM task_assignees WHERE task_id = $1 AND user_id = $2`
	row := db.QueryRowContext(ctx, sqlstr, taskID, userID)

	ta := xo.TaskAssignee{}
	if err := row.Scan(&ta.ID, &ta.TaskID, &ta.UserID, &ta.CreatedAt, &ta.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Newf(codes.CodeNotFound, "task assignee not found")
		}
		return nil, errors.Convert(err)
	}

	return &entity.TaskAssignee{
		ID:     ta.ID,
		TaskID: ta.TaskID,
		UserID: ta.UserID,
	}, nil
}
