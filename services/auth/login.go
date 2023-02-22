package services

import (
	"echo-mongo-api/models"
	"echo-mongo-api/responses"
	services "echo-mongo-api/services/user"
	"echo-mongo-api/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	userFounded, err := services.FindUserByEmail(user.Email)

	if err != nil {
		return echo.ErrUnauthorized
	}

	claims := jwt.MapClaims{}
	claims["name"] = userFounded.Name
	claims["email"] = userFounded.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		return errToken
	}

	return c.JSON(http.StatusOK, responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &echo.Map{
			"token": t,
		},
	})
}
