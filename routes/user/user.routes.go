package user_route

import (
	"echo-mongo-api/controllers"
	"github.com/labstack/echo/v4"
)

func SetupUserRoute(e *echo.Group) {
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetOneUser)
	e.PUT("/user/:userId", controllers.EditUser)
	e.DELETE("/user/:userId", controllers.DeleteUser)
	e.GET("/user", controllers.GetAllUsers)
}
