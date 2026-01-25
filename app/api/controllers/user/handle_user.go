package user

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *UserController) Me(c *gin.Context) {
	identity, err := s.userService.Identity(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	utils.Success(c, identity)
}
func (s *UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	_ = session.Save()
	utils.Success(c, nil, "logged out")
}
