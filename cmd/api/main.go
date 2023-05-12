package main

import (
	"os"

	"github.com/ilhamnyto/echo-fw/config"
	"github.com/ilhamnyto/echo-fw/controller"
	"github.com/ilhamnyto/echo-fw/repositories"
	"github.com/ilhamnyto/echo-fw/routes"
	"github.com/ilhamnyto/echo-fw/services"

	"github.com/ilhamnyto/echo-fw/pkg/database"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig(".env")

	db := database.ConnectDB()

	e := echo.New()

	userRepository := repositories.NewUserRepository(db.DbSQL)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	routes.UserRouter(e, userController)


	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
	
}