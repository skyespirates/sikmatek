package pgsql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type taskRepo struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepo{
		db: db,
	}
}

func (tp *taskRepo) GetAll(ctx context.Context) ([]*entity.Task, error) {

	user := utils.ContextGetUser(ctx)

	query := `SELECT id, title, is_completed, created_at, updated_at FROM tasks WHERE user_id = $1 ORDER BY updated_at DESC NULLS LAST, created_at DESC`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := tp.db.QueryContext(ctx, query, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task

	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.Id, &task.Title, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tp *taskRepo) GetById(ctx context.Context, id int) (*entity.Task, error) {
	query := `SELECT id, title, is_completed, created_at, updated_at FROM tasks WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var task entity.Task

	err := tp.db.QueryRowContext(ctx, query, id).Scan(&task.Id, &task.Title, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (tp *taskRepo) Create(ctx context.Context, title string) (*entity.Task, error) {

	user := utils.ContextGetUser(ctx)

	query := `INSERT INTO tasks (title, user_id) VALUES ($1, $2) RETURNING id, title, is_completed, created_at, updated_at`
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	args := []any{title, user.Id}

	var task entity.Task
	err := tp.db.QueryRowContext(ctx, query, args...).Scan(&task.Id, &task.Title, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (tp *taskRepo) Delete(ctx context.Context, id int) (int, error) {
	query := `DELETE FROM tasks WHERE id = $1 RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var todoId int

	err := tp.db.QueryRowContext(ctx, query, id).Scan(&todoId)
	if err != nil {
		return 0, err
	}
	return todoId, nil
}

func (tp *taskRepo) Update(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	user := utils.ContextGetUser(ctx)
	query := `UPDATE tasks SET title = $1, is_completed = $2, updated_at = NOW() WHERE id = $3 AND user_id = $4 RETURNING id, title, is_completed, created_at, updated_at`
	args := []any{task.Title, task.IsCompleted, task.Id, user.Id}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tp.db.QueryRowContext(ctx, query, args...).Scan(&task.Id, &task.Title, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil, err
	}
	return task, nil
}
