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

	port := config.Int("port")

	if err := router.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v\n", err))
	}
}
