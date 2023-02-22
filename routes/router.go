package routes

import (
	userRoutes "echo-mongo-api/routes/user"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	api := app.Group("/api")

	userRoutes.SetupUserRoute(api)
	TransactionRoute(api)
}
