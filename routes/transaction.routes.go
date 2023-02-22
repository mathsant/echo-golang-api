package routes

import (
	"echo-mongo-api/controllers"
	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Group) {
	e.POST("/transaction", controllers.CreateTransactionByUser)
}
