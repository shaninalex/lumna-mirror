package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *WorkspaceController) handlerList(c *gin.Context) {
	workspaces, err := s.workspaceManager.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err)
		return
	}

	utils.Success(c, workspaces)
}
