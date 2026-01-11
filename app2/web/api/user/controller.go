package user

import (
	"gitlab.com/shaninalex/lumna/app2/web"
)

func RegisterUserController(router *web.Router) {
	h := NewUserHandler()

	// fetch current user info
	router.GET("/api/v1/user/me", h.HandleGetUser)

	// update user settings
	// router.POST("/api/v1/user/settings", h.HandleUpdateSettings)

	// logout user
	router.GET("/api/v1/user/logout", h.HandleLogout)
}
