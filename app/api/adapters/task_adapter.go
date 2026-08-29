package adapters

import (
	"gitlab.com/shaninalex/lumna/app/api/dto"
	"gitlab.com/shaninalex/lumna/app/models"
)

func ToTaskDto(task models.Task) dto.TaskDto {
	return dto.TaskDto{
		ID:           task.ID,
		Title:        task.Title,
		Body:         task.Body,
		Completed:    task.Completed,
		Meta:         task.Meta,
		ProjectId:    task.ProjectId,
		Boards:       dto.ToBoardTaskDtoList(task.Boards),
		OwnerId:      task.OwnerId,
		AssigneesIDs: task.AssigneesIDs,
		CreatedAt:    task.CreatedAt,
		UpdatedAt:    task.UpdatedAt,
		TaskEvents:   []dto.EntityEventDTO{},
	}
}

func ToTaskDtoList(tasks []models.Task, events []models.EntityEvent) []dto.TaskDto {
	result := make([]dto.TaskDto, 0, len(tasks))
	for _, task := range tasks {
		td := ToTaskDto(task)

		for _, e := range events {
			if e.EntityId == nil || e.EntityType == nil {
				continue
			}

			if task.ID == *e.EntityId && *e.EntityType == "task" {
				td.TaskEvents = append(td.TaskEvents, dto.ToEntityEventDTO(e))
			}
		}

		result = append(result, td)
	}
	return result
}
