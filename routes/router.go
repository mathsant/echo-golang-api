package routes

import (
	userRoutes "echo-mongo-api/routes/user.route"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	api := app.Group("/api")

	userRoutes.SetupUserRoute(api)
}
