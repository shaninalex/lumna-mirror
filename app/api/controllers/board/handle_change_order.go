package board

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

type KTask struct {
	Id    string `json:"id"`
	Order int    `json:"order"`
}

type KColumn struct {
	Id    string  `json:"id"`
	Order *int    `json:"order"`
	Tasks []KTask `json:"tasks"`
}

type KanbanBoardChangeOrderPayload struct {
	MoveType string    `json:"moveType"`
	ColumnId *string   `json:"columnId"`
	Tasks    []KTask   `json:"tasks"`
	Columns  []KColumn `json:"columns"`
}

func (s *BoardController) ChangeOrder(c *gin.Context) {
	var payload KanbanBoardChangeOrderPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	_db := db.GetDB(ctx)

	if payload.MoveType == "task" {
		// move in a single column
		if payload.ColumnId != nil {
			for _, task := range payload.Tasks {
				if result := _db.Model(&models.Task{}).Where("id = ?", task.Id).Update("order", task.Order); result.Error != nil {
					utils.Error(c, http.StatusBadRequest, result.Error)
					return
				}
			}
		} else {
			// move task between columns
			for _, column := range payload.Columns {
				for _, task := range column.Tasks {
					result := _db.Model(&models.Task{}).
						Where("id = ?", task.Id).
						Updates(map[string]any{
							"order":     task.Order,
							"column_id": column.Id,
						})
					if result.Error != nil {
						utils.Error(c, http.StatusBadRequest, result.Error)
						return
					}
				}
			}
		}
	} else {
		// move column between each other
		for _, column := range payload.Columns {
			if result := _db.Model(&models.Column{}).Where("id = ?", column.Id).Update("order", column.Order); result.Error != nil {
				utils.Error(c, http.StatusBadRequest, result.Error)
				return
			}
		}
	}

	utils.Success(c, nil)
}
