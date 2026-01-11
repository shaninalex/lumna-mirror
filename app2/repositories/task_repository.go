package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app2/models"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
)

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

var _ Repository[models.Task] = (*TaskRepository)(nil)

func (s *TaskRepository) Get(ctx context.Context, id uint) (*models.Task, error) {
	task := &models.Task{}
	row := db.FromContext(ctx).QueryRow(`
		SELECT id, board_id, list_id, name, done, list_order, created_at, updated_at FROM tasks WHERE id = ?
	`, id)
	err := row.Scan(
		&task.Id,
		&task.BoardId,
		&task.ListId,
		&task.Name,
		&task.Done,
		&task.Order,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorProjectNotFound
		}
		return nil, err
	}

	return task, nil
}

func (s *TaskRepository) Delete(ctx context.Context, id uint) error {
	res, err := db.FromContext(ctx).ExecContext(ctx, `DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}

func (s *TaskRepository) Create(ctx context.Context, entry *models.Task) error {
	query := `
		INSERT INTO tasks (board_id, list_id, name, list_order, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	entry.SetCreatedAt(time.Now())
	entry.SetUpdatedAt(time.Now())

	result, err := db.FromContext(ctx).ExecContext(
		ctx, query,
		entry.BoardId,
		entry.ListId,
		entry.Name,
		entry.Order,
		entry.CreatedAt,
		entry.UpdatedAt,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	entry.SetId(uint(id))

	return nil
}

func (s *TaskRepository) List(ctx context.Context, opts db.Expr) ([]*models.Task, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT id, board_id, list_id, name, list_order, done, created_at, updated_at FROM tasks %s`,
		where,
	)

	rows, err := db.FromContext(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*models.Task{}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.Id,
			&task.BoardId,
			&task.ListId,
			&task.Name,
			&task.Order,
			&task.Done,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (s *TaskRepository) Count(ctx context.Context, opts db.Expr) (int, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT count(*) FROM tasks %s`,
		where,
	)

	var count int
	row := db.FromContext(ctx).QueryRowContext(ctx, query, args...)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (s *TaskRepository) Update(ctx context.Context, taskId uint, opts db.SetExpr) error {
	setQuery, args := opts.Build()
	query := fmt.Sprintf(
		`UPDATE tasks %s WHERE id = ?`,
		setQuery,
	)

	args = append(args, taskId)

	res, err := db.FromContext(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}
