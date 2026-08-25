package adapters

import (
	"slices"

	"gitlab.com/shaninalex/lumna/app/api/dto"
	"gitlab.com/shaninalex/lumna/app/models"
)

func ToTaskDto(task models.Task, boardTask models.BoardTask) dto.TaskDto {
	return dto.TaskDto{
		ID:        task.ID,
		Title:     task.Title,
		Body:      task.Body,
		Completed: task.Completed,
		Meta:      task.Meta,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		Board:     dto.ToBoardTaskDto(boardTask),
	}
}

func ToTaskDtoList(tasks []models.Task, boardTasks []models.BoardTask) []dto.TaskDto {
	result := make([]dto.TaskDto, 0, len(tasks))
	for _, task := range tasks {
		td := dto.TaskDto{
			ID:        task.ID,
			Title:     task.Title,
			Body:      task.Body,
			Completed: task.Completed,
			Meta:      task.Meta,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
		i := slices.IndexFunc(boardTasks, func(bt models.BoardTask) bool { return bt.TaskId == task.ID })
		if i >= 0 {
			td.Board = dto.ToBoardTaskDto(boardTasks[i])
		}
		result = append(result, td)
	}
	return result
}
