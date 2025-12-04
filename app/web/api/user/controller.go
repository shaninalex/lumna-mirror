package user

import (
	"fmt"

	"gitlab.com/shaninalex/lumna/app/web"
)

func RegisterUserController(router *web.Router, baseUrl string) {
	h := NewUserHandler()

	// fetch current user info
	router.GET(fmt.Sprintf("%s/me", baseUrl), h.HandleGetUser)

	// update user settings
	router.POST(fmt.Sprintf("%s/settings", baseUrl), h.HandleUpdateSettings)

	// logout user
	router.GET(fmt.Sprintf("%s/logout", baseUrl), h.HandleLogout)
}
