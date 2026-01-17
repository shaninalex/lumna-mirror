package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

func NewController(router *gin.RouterGroup) {
	controller := NewUserContoller()

	router.GET("/api/v1/user/me", controller.Me)
	router.GET("/api/v1/user/logout", controller.Logout)
}

type UserController struct {
	userService *services.UserService
}

func NewUserContoller() *UserController {
	s := &UserController{
		userService: &services.UserService{},
	}

	return s
}
