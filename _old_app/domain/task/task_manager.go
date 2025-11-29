package task

import (
	"context"
	"log"
	"time"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/db"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/utils"
)

// TaskReader - task reader.
type TaskReader interface {
	TasksList(ctx context.Context, projectID int64) ([]*Task, error)
	TaskDetail(ctx context.Context, taskID int64) (*Task, error)
}

// TaskWriter - task writer.
type TaskWriter interface {
	TaskUpdate(ctx context.Context, data *Task) error
	TaskCreate(ctx context.Context, data *Task) (*Task, error)
	TaskDelete(ctx context.Context, taskID int64) error
}

type TaskManager interface {
	TaskReader
	TaskWriter
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

type TaskService struct{}

func (t TaskService) TaskDelete(ctx context.Context, taskID int64) error {
	return TaskDelete(ctx, db.GetDb(ctx), taskID)
}

func (t TaskService) TasksList(ctx context.Context, projectID int64) ([]*Task, error) {
	tasks, err := TaskList(ctx, db.GetDb(ctx), projectID)
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if err = getCommentsCount(ctx, task); err != nil {
			continue
		}
	}
	return tasks, nil
}

func (t TaskService) TaskDetail(ctx context.Context, taskID int64) (*Task, error) {
	task, err := TaskGet(ctx, db.GetDb(ctx), taskID)
	if err != nil {
		return nil, err
	}
	getTaskRelations(ctx, task)
	return task, nil
}

func (t TaskService) TaskUpdate(ctx context.Context, data *Task) error {
	return TaskUpdate(ctx, db.GetDb(ctx), data.Id, &Task{
		Title:       data.Title,
		StatusID:    data.StatusID,
		UserID:      data.UserID,
		Description: data.Description,
		Completed:   data.Completed,
		ListIndex:   data.ListIndex,
	})
}

func (t TaskService) TaskCreate(ctx context.Context, data *Task) (*Task, error) {
	maxIndex := TaskGetIndex(ctx, db.GetDb(ctx), data.StatusID)
	task := Task{
		UserID:      data.UserID,
		ProjectID:   data.ProjectID,
		StatusID:    data.StatusID,
		Title:       data.Title,
		Code:        utils.GenerateEntityCode("task"),
		Completed:   data.Completed,
		Description: data.Description,
		ListIndex:   maxIndex,
	}
	err := TaskSave(ctx, db.GetDb(ctx), &task)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	data.Id = task.Id
	data.CreatedAt = now
	data.UpdatedAt = now
	data.Code = task.Code
	data.ListIndex = task.ListIndex
	return data, nil
}

func getTaskRelations(ctx context.Context, task *Task) {
	if err := getBadges(ctx, task); err != nil {
		log.Println(err)
	}
	if err := getComments(ctx, task); err != nil {
		log.Println(err)
	}
}

func getComments(ctx context.Context, task *Task) error {
	comments, err := CommentsList(ctx, db.GetDb(ctx), task.Id, EntityTypeTask)
	if err != nil {
		return err
	}
	for _, comment := range comments {
		task.Comments = append(task.Comments, &Comment{
			Id:         comment.Id,
			EntityId:   comment.EntityId,
			EntityType: comment.EntityType,
			UserId:     comment.UserId,
			Content:    comment.Content,
			CreatedAt:  comment.CreatedAt,
		})
	}

	task.CommentsCount = len(comments)

	return nil
}

func getCommentsCount(ctx context.Context, task *Task) error {
	count, err := CommentsCount(ctx, db.GetDb(ctx), task.Id, EntityTypeTask)
	if err != nil {
		return err
	}
	task.CommentsCount = count
	return nil
}

func getBadges(ctx context.Context, task *Task) error {
	badges, err := BadgeTaskList(ctx, db.GetDb(ctx), task.Id)
	if err != nil {
		return err
	}
	task.Badges = badges
	return nil
}
