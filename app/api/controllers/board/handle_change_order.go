package board

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
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

/*

Change conditions:
- Task:
in a single column:
columnId = nil
tasks = list of 2 different tasks with orders
columns = nil

move to another column:
columnId = nil
tasls = empty
columns = contain 2 columns with all tasks { id, order }

- Column:
ColumnId = nil
tasks = empty
columns = contain all columns of boards { id, order }

*/

func (s *BoardController) ChangeOrder(c *gin.Context) {
	var payload KanbanBoardChangeOrderPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
	}

	log.Println(payload)
	if payload.MoveType == "task" {

	}

	if payload.MoveType == "column" {

	}

	// TODO: handle change
	// s.taskService.ReorderTask()
	// s.columnService.Reorder()

	utils.Success(c, nil)
}
