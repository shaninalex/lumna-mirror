package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/internal/db"
)

type Comment struct {
	Id        int64
	TaskId    int64
	UserId    int64
	UserName  string
	UserImage string
	Content   string
	CreatedAt time.Time
}

type CommentManager interface {
	List(ctx context.Context, taskId int64) ([]*Comment, error)
	CreateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id int64) error
}

func NewCommentService() CommentManager {
	return &CommentService{}
}

type CommentService struct{}

func (c CommentService) List(ctx context.Context, taskId int64) ([]*Comment, error) {
	dbComments, err := db.CommentsList(ctx, db.GetDb(ctx), taskId)
	if err != nil {
		return nil, err
	}
	comments := make([]*Comment, len(dbComments))
	for i, comment := range dbComments {
		comments[i] = ToDomainComment(comment)
	}
	return comments, nil
}

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

func ToDomainComment(comment *db.Comment) *Comment {
	return &Comment{
		Id:        comment.ID,
		TaskId:    comment.TaskID,
		UserId:    comment.UserID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
}
