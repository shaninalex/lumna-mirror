package comment

import (
	"context"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/db"
)

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
