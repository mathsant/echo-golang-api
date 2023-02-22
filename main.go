package main

import (
	"echo-mongo-api/configs"
	"echo-mongo-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}
