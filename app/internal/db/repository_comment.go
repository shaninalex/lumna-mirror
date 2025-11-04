package db

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// CommentGet get comment by id
func CommentGet(ctx context.Context, db *sql.DB, id int64) (*Comment, error) {
	row := db.QueryRowContext(ctx, `
		SELECT id, task_id, user_id, content, created_at
		FROM comments
		WHERE id = ?`, id)

	var b Comment
	if err := row.Scan(&b.ID, &b.TaskID, &b.UserID, &b.Content, &b.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // not found
		}
		return nil, err
	}
	return &b, nil
}

// CommentCreate inserts a new comment and updates its ID.
func CommentCreate(ctx context.Context, db *sql.DB, comment *Comment) error {
	res, err := db.ExecContext(ctx, `
		INSERT INTO comments (task_id, user_id, content)
		VALUES (?, ?, ?)`,
		comment.TaskID, comment.UserID, comment.Content)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	comment.ID = id
	comment.CreatedAt = time.Now()
	return nil
}

func CommentDelete(ctx context.Context, db *sql.DB, id int64) error {
	_, err := db.ExecContext(ctx, `DELETE FROM comments WHERE id = ?`, id)
	return err
}

// CommentsList returns all comments belonging to a task.
func CommentsList(ctx context.Context, db *sql.DB, taskID int64) ([]*Comment, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id, task_id, user_id, content, created_at
		FROM comments
		WHERE project_id = ?
		ORDER BY created_at ASC`, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Comment
	for rows.Next() {
		var b Comment
		if err := rows.Scan(&b.ID, &b.TaskID, &b.UserID, &b.Content, &b.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &b)
	}
	return result, rows.Err()
}

