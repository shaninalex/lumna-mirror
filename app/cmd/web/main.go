package main

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/shaninalex/lumna/app/base"
	"gitlab.com/shaninalex/lumna/app/web"
)

func main() {
	fmt.Println("Run Lumna as a webserver")

	config := base.GetConfig()

	router := web.NewDefaultRouter()

	// TODO: make controller registry
	// For example:
	// router.RegisterController("/auth", auth.NewAuthController)
	// router.RegisterController("/user", auth.NewUserController)
	// router.RegisterController("/projects", auth.NewProjectsController)
	// router.RegisterController("/task", auth.NewTaskController)

	port := config.Int("port")

	if err := router.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v\n", err))
	}
}
