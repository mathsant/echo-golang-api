package controllers

import (
	"echo-mongo-api/models"
	"echo-mongo-api/responses"
	services "echo-mongo-api/services/transaction"
	"echo-mongo-api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateTransactionByUser(c echo.Context) error {
	var userId string
	var transaction models.Transaction

	tokenClaims, err := utils.GetTokensClaims(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	for key, value := range tokenClaims {
		if key == "userId" {
			userId = value.(string)
		}
	}

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	transactionCreated, err := services.CreateTransactionByUser(transaction, userId)

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

func GetAllTransactionsByUser(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	limitConverted, _ := strconv.Atoi(limit)
	pageConverted, _ := strconv.Atoi(page)

	var userId string

	tokenClaims, err := utils.GetTokensClaims(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	for key, value := range tokenClaims {
		if key == "userId" {
			userId = value.(string)
		}
	}

	transactions, err := services.GetTransactionsByUser(userId, limitConverted, pageConverted)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{
		Status:  http.StatusOK,
		Message: "error",
		Data:    &echo.Map{"data": transactions},
	})
}

func GetTransactionByDate(c echo.Context) error {
	var userId string

	tokenClaims, err := utils.GetTokensClaims(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	for key, value := range tokenClaims {
		if key == "userId" {
			userId = value.(string)
		}
	}

	transactions, err := services.GetTransactionsByDate(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{
		Status:  http.StatusOK,
		Message: "error",
		Data:    &echo.Map{"data": transactions},
	})
}
