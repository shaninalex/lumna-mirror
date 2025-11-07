package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/internal/db"
)

type Comment struct {
	Id        int64     `json:"id" db:"id"`
	TaskId    int64     `json:"task_id" db:"user_id"`
	UserId    int64     `json:"user_id" db:"project_id"`
	Content   string    `json:"content" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CommentManager interface {
	CreateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id int64) error
}

func NewCommentService() CommentManager {
	return &CommentService{}
}

type CommentService struct{}

func (c CommentService) CreateComment(ctx context.Context, comment *Comment) error {
	dbComment := &db.Comment{
		TaskID:    comment.TaskId,
		UserID:    comment.UserId,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
	if err := db.CommentCreate(ctx, db.GetDb(ctx), dbComment); err != nil {
		return err
	}

	comment.Id = dbComment.ID
	comment.CreatedAt = dbComment.CreatedAt
	return nil
}

func (c CommentService) DeleteComment(ctx context.Context, id int64) error {
	if err := db.CommentDelete(ctx, db.GetDb(ctx), id); err != nil {
		return err
	}
	return nil
}
