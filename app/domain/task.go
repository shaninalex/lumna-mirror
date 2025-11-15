package domain

import (
	"context"
	"log"
	"time"

	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/internal/utils"
)

const EntityTypeTask = "task"

type Task struct {
	ID          int64
	UserID      int64
	ProjectID   int64
	StatusID    int64
	Title       string
	Completed   bool
	Description *string
	ListIndex   float64
	Code        string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// related structures
	Badges        []*Badge
	Comments      []*Comment
	CommentsCount int
}

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
	return db.TaskDelete(ctx, db.GetDb(ctx), taskID)
}

func (t TaskService) TasksList(ctx context.Context, projectID int64) ([]*Task, error) {
	dbTasks, err := db.TaskList(ctx, db.GetDb(ctx), projectID)
	if err != nil {
		return nil, err
	}
	tasks := make([]*Task, len(dbTasks))
	for i, task := range dbTasks {
		tasks[i] = &Task{
			ID:          task.ID,
			UserID:      task.UserID,
			ProjectID:   task.ProjectID,
			StatusID:    task.StatusID,
			Title:       task.Title,
			Completed:   task.Completed,
			Description: task.Description,
			ListIndex:   task.ListIndex,
			Code:        task.Code,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}
		if err = getCommentsCount(ctx, tasks[i]); err != nil {
			continue
		}
	}
	return tasks, nil
}

func (t TaskService) TaskDetail(ctx context.Context, taskID int64) (*Task, error) {
	task, err := db.TaskGet(ctx, db.GetDb(ctx), taskID)
	if err != nil {
		return nil, err
	}
	model := &Task{
		ID:          task.ID,
		UserID:      task.UserID,
		ProjectID:   task.ProjectID,
		StatusID:    task.StatusID,
		Title:       task.Title,
		Completed:   task.Completed,
		Description: task.Description,
		ListIndex:   task.ListIndex,
		Code:        task.Code,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	getTaskRelations(ctx, model)
	return model, nil
}

func (t TaskService) TaskUpdate(ctx context.Context, data *Task) error {
	return db.TaskUpdate(ctx, db.GetDb(ctx), data.ID, &db.Task{
		Title:       data.Title,
		StatusID:    data.StatusID,
		UserID:      data.UserID,
		Description: data.Description,
		Completed:   data.Completed,
		ListIndex:   data.ListIndex,
	})
}

func (t TaskService) TaskCreate(ctx context.Context, data *Task) (*Task, error) {
	maxIndex := db.TaskGetIndex(ctx, db.GetDb(ctx), data.StatusID)
	task := db.Task{
		UserID:      data.UserID,
		ProjectID:   data.ProjectID,
		StatusID:    data.StatusID,
		Title:       data.Title,
		Code:        utils.GenerateEntityCode("task"),
		Completed:   data.Completed,
		Description: data.Description,
		ListIndex:   maxIndex,
	}
	err := db.TaskSave(ctx, db.GetDb(ctx), &task)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	data.ID = task.ID
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
	comments, err := CommentsList(ctx, db.GetDb(ctx), task.ID, EntityTypeTask)
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
	count, err := CommentsCount(ctx, db.GetDb(ctx), task.ID, EntityTypeTask)
	if err != nil {
		return err
	}
	task.CommentsCount = count
	return nil
}

func getBadges(ctx context.Context, task *Task) error {
	badges, err := BadgeTaskList(ctx, db.GetDb(ctx), task.ID)
	if err != nil {
		return err
	}
	task.Badges = badges
	return nil
}
