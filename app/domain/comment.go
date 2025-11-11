package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/internal/db"
)

type Comment struct {
	Id         int64     `json:"id" db:"id"`
	EntityId   int64     `json:"entity_id" db:"entity_id"`
	EntityType string    `json:"entity_type" db:"entity_type"`
	UserId     int64     `json:"user_id" db:"project_id"`
	Content    string    `json:"content" db:"title"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type CommentManager interface {
	ListComments(ctx context.Context, entityId int64, entityType string) []*Comment
	CreateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id int64) error
}

func NewCommentService() CommentManager {
	return &CommentService{}
}

type CommentService struct{}

func (c CommentService) ListComments(ctx context.Context, entityId int64, entityType string) []*Comment {
	dbComments, err := db.CommentsList(ctx, db.GetDb(ctx), entityId, entityType)
	if err != nil {
		return []*Comment{}
	}

	comments := make([]*Comment, len(dbComments))
	for i, c := range dbComments {
		comments[i] = &Comment{
			Id:         c.ID,
			EntityId:   c.EntityId,
			EntityType: c.EntityType,
			UserId:     c.UserID,
			Content:    c.Content,
			CreatedAt:  c.CreatedAt,
		}
	}

	return comments
}

func (c CommentService) CreateComment(ctx context.Context, comment *Comment) error {
	dbComment := &db.Comment{
		EntityId:   comment.EntityId,
		EntityType: comment.EntityType,
		UserID:     comment.UserId,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt,
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
