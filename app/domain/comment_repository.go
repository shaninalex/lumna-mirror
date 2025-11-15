package domain

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// CommentGet get comment by id
func CommentGet(ctx context.Context, db *sql.DB, id int64) (*Comment, error) {
	row := db.QueryRowContext(ctx, `
		SELECT id, entity_id, entity_type, user_id, content, created_at
		FROM comments
		WHERE id = ?`, id)

	var b Comment
	if err := row.Scan(&b.Id, &b.EntityId, &b.EntityType, &b.UserId, &b.Content, &b.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // not found
		}
		return nil, err
	}
	return &b, nil
}

// CommentCreate inserts a new comment and updates its Id.
func CommentCreate(ctx context.Context, db *sql.DB, comment *Comment) error {
	res, err := db.ExecContext(ctx, `
		INSERT INTO comments (entity_id, entity_type, user_id, content)
		VALUES (?, ?, ?, ?)`,
		comment.EntityId, comment.EntityType, comment.UserId, comment.Content)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	comment.Id = id
	comment.CreatedAt = time.Now()
	return nil
}

func CommentDelete(ctx context.Context, db *sql.DB, id int64) error {
	_, err := db.ExecContext(ctx, `DELETE FROM comments WHERE id = ?`, id)
	return err
}

// CommentsList returns all comments belonging to a task.
func CommentsList(ctx context.Context, db *sql.DB, entityId int64, entityType string) ([]*Comment, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id, entity_id, entity_type, user_id, content, created_at
		FROM comments
		WHERE entity_id = ? and entity_type = ?
		ORDER BY created_at ASC`, entityId, entityType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Comment
	for rows.Next() {
		var b Comment
		if err := rows.Scan(&b.Id, &b.EntityId, &b.EntityType, &b.UserId, &b.Content, &b.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &b)
	}
	return result, rows.Err()
}

// CommentsCount returns amount of comments for a task.
func CommentsCount(ctx context.Context, db *sql.DB, entityId int64, entityType string) (int, error) {
	row := db.QueryRowContext(ctx, `
		SELECT count(*)
		FROM comments
		WHERE entity_id = ? and entity_type = ?
		ORDER BY created_at ASC`, entityId, entityType)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
