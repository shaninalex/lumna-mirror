package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

func RegisterUserController(router *gin.RouterGroup) {
	controller := NewUserContoller()

	router.GET("/user/me", controller.Me)
	router.GET("/user/logout", controller.Logout)
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
