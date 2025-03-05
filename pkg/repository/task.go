package repository

import (
	"context"
	"database/sql"
	"time"

	xo "github.com/yudai2929/task-app/database/gen"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) GetTask(ctx context.Context, id string) (*entity.Task, error) {
	t, err := xo.TaskByID(ctx, r.db, id)
	if err != nil {
		return nil, errors.Convert(err)
	}

	return r.convertTask(t), nil
}

func (r *taskRepository) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	now := time.Now()

	t := &xo.Task{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      int(task.Status),
		DueDate:     convertTime(task.DueDate),
		CreatedAt: sql.NullTime{
			Time:  now,
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  now,
			Valid: true,
		},
	}
	if err := t.Insert(ctx, r.db); err != nil {
		return nil, errors.Convert(err)
	}
	return r.convertTask(t), nil
}

func (r *taskRepository) ListTasks(ctx context.Context) (entity.Tasks, error) {
	const sqlstr = `SELECT id, user_id, title, description, status, due_date, created_at, updated_at ` +
		`FROM public.tasks ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, errors.Convert(err)
	}
	defer rows.Close()

	var tasks entity.Tasks
	for rows.Next() {
		t := xo.Task{}
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, errors.Convert(err)
		}
		tasks = append(tasks, r.convertTask(&t))
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Convert(err)
	}
	return tasks, nil
}

func (r *taskRepository) ListMyTasks(ctx context.Context, userID string) (entity.Tasks, error) {
	const sqlstr = `
		SELECT DISTINCT t.id, t.user_id, t.title, t.description, t.status, t.due_date, t.created_at, t.updated_at
		FROM tasks t
		LEFT JOIN task_assignees ta ON t.id = ta.task_id
		WHERE t.user_id = $1 OR ta.user_id = $1
		ORDER BY t.created_at DESC`
	rows, err := r.db.QueryContext(ctx, sqlstr, userID)
	if err != nil {
		return nil, errors.Convert(err)
	}
	defer rows.Close()

	var tasks entity.Tasks
	for rows.Next() {
		t := xo.Task{}
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, errors.Convert(err)
		}
		tasks = append(tasks, r.convertTask(&t))
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Convert(err)
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	var updatedTask *entity.Task
	err := runInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		t, err := xo.TaskByID(ctx, tx, task.ID)
		if err != nil {
			return errors.Convert(err)
		}

		t.UserID = task.UserID
		t.Title = task.Title
		t.Description = task.Description
		t.Status = int(task.Status)
		t.DueDate = convertTime(task.DueDate)
		t.UpdatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}

		if err := t.Update(ctx, tx); err != nil {
			return errors.Convert(err)
		}
		updatedTask = r.convertTask(t)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id string) error {
	return runInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		t, err := xo.TaskByID(ctx, tx, id)
		if err != nil {
			return errors.Convert(err)
		}

		if err := t.Delete(ctx, tx); err != nil {
			return errors.Convert(err)
		}
		return nil
	})
}

func (r *taskRepository) convertTask(t *xo.Task) *entity.Task {
	return &entity.Task{
		ID:          t.ID,
		UserID:      t.UserID,
		Title:       t.Title,
		Description: t.Description,
		Status:      entity.TaskStatus(t.Status),
		DueDate:     convertNullTime(t.DueDate),
	}
}

func convertNullTime(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func convertTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{
			Time:  *t,
			Valid: true,
		}
	}
	return sql.NullTime{
		Valid: false,
	}
}
