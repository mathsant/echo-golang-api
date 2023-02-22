package controllers

import (
	"echo-mongo-api/models"
	"echo-mongo-api/responses"
	services "echo-mongo-api/services/transaction"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTransactionByUser(c echo.Context) error {
	var transaction models.Transaction

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	transactionCreated, err := services.CreateTransactionByUser(transaction, transaction.UserId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &echo.Map{"data": transactionCreated},
	})
}
