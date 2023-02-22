package user_route

import (
	"echo-mongo-api/controllers"
	"echo-mongo-api/middleware"
	services "echo-mongo-api/services/auth"
	"github.com/labstack/echo/v4"
)

func SetupUserRoute(e *echo.Group) {
	e.POST("/user", controllers.CreateUser, middleware.IsLoggedIn)
	e.GET("/user/:userId", controllers.GetOneUser, middleware.IsLoggedIn)
	e.PUT("/user/:userId", controllers.EditUser, middleware.IsLoggedIn)
	e.DELETE("/user/:userId", controllers.DeleteUser, middleware.IsLoggedIn)
	e.GET("/user", controllers.GetAllUsers, middleware.IsLoggedIn)
	e.POST("/login", services.Login)
}
