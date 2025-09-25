// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package main

func main() {
	//args := os.Args
	//var configPath string
	//if len(args) < 2 || os.Getenv("CONFIG_PATH") != "" {
	//	configPath = os.Getenv("CONFIG_PATH")
	//} else {
	//	configPath = args[1]
	//}
	//
	//config := base.NewConfig(configPath)
	//db := database.InitDB(database.BuildDSN(config))
	//sqlDB, err := db.DB()
	//if err != nil {
	//	panic(err)
	//}
	//router := web.AuthRouter(config, sqlDB, "user_service")
	//userApp.NewUserController(router)
	//
	//srv := &http.Server{
	//	Handler:      router,
	//	Addr:         fmt.Sprintf(":%s", config.String("user_service.port")),
	//	ReadTimeout:  15 * time.Second,
	//	WriteTimeout: 30 * time.Second,
	//}
	//if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//	log.Printf("Server error: %v\n", err)
	//}
}
