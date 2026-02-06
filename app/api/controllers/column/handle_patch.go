package column

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *ColumnController) Patch(c *gin.Context) {
	columnId, err := uuid.Parse(c.Param("columnId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	payload := services.ColumnUpdate{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	var column models.Column
	if result := persistence.GetDB(c.Request.Context()).Where("id = ?", columnId).First(&column); result.Error != nil {
		log.Println(result.Error)
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	column.Title = payload.Title

	if _, err := s.columnService.Update(c.Request.Context(), &column); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, column)
}
