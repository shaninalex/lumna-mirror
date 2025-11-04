package domain

import (
	"context"
	"time"
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
	CreateComment(ctx context.Context, project *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, id int64) error
}
