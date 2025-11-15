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

// GetID - returns the id.
func (s *Comment) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *Comment) SetID(id int64) { s.Id = id }

// GetOwnerID - returns the owner id.
func (s *Comment) GetOwnerID() int64 { return s.UserId }

// IsOwner - checks if it is owner.
func (s *Comment) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Comment) GetCreatedAt() time.Time { return s.CreatedAt }

// GetCreatedBy - returns the created by.
func (s *Comment) GetCreatedBy() int64 { return s.UserId }

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
	comments, err := CommentsList(ctx, db.GetDb(ctx), entityId, entityType)
	if err != nil {
		return []*Comment{}
	}
	return comments
}

func (c CommentService) CreateComment(ctx context.Context, comment *Comment) error {
	if err := CommentCreate(ctx, db.GetDb(ctx), comment); err != nil {
		return err
	}
	return nil
}

func (c CommentService) DeleteComment(ctx context.Context, id int64) error {
	if err := CommentDelete(ctx, db.GetDb(ctx), id); err != nil {
		return err
	}
	return nil
}
