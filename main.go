package main

import (
	"echo-mongo-api/configs"
	"echo-mongo-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	configs.ConnectDB()

	routes.SetupRoutes(app)

	app.Logger.Fatal(app.Start(":8000"))
}
