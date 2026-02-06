package board

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *BoardController) Patch(c *gin.Context) {
	boardId, err := uuid.Parse(c.Param("boardId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	payload := struct {
		ProjectId uuid.UUID `json:"project_id"`
		Title     string    `json:"title"`
	}{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	// use board service
	var board models.Board
	if result := persistence.GetDB(c.Request.Context()).Where("id = ?", boardId).First(&board); result.Error != nil {
		log.Println(result.Error)
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	board.Title = payload.Title

	if result := persistence.GetDB(c.Request.Context()).Save(&board); result.Error != nil {
		log.Println(result.Error)
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, board)
}
