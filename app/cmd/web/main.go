package main

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/shaninalex/lumna/app/base"
	"gitlab.com/shaninalex/lumna/app/web"
)

func main() {
	config := base.GetConfig()
	fmt.Println("Run Lumna as a webserver")

	router := web.NewDefaultRouter()

	static := web.GetStaticFS()

	// Public controllers
	if static != nil {
		router.GET("/", web.FrontendHandler(static))
	}

	port := config.Int("port")
	if err := router.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v", err))
	}
}
