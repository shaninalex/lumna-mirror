package adapter

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/domain"
)

type CommentDto struct {
	Id        int64     `json:"id"`
	TaskId    int64     `json:"task_id"`
	UserId    int64     `json:"user_id"`
	UserName  string    `json:"user_name"`
	UserImage string    `json:"user_image"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCommentDto(c *domain.Comment) *CommentDto {
	return &CommentDto{
		Id:        c.Id,
		TaskId:    c.TaskId,
		UserId:    c.UserId,
		UserName:  c.UserName,
		UserImage: c.UserImage,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}

func NewCommentsDtoList(comments []*domain.Comment) []*CommentDto {
	var dtos []*CommentDto
	for _, comment := range comments {
		dtos = append(dtos, NewCommentDto(comment))
	}
	return dtos
}

func NewDomainCommentFromDto(c *CommentDto) *domain.Comment {
	return &domain.Comment{
		Id:        c.Id,
		TaskId:    c.TaskId,
		UserId:    c.UserId,
		UserName:  c.UserName,
		UserImage: c.UserImage,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}
