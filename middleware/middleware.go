package middleware

import (
	"echo-mongo-api/responses"
	"echo-mongo-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("SECRET_TOKEN"),
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var isAdmin bool
		tokenClaims, err := utils.GetTokensClaims(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, responses.UserResponse{
				Status:  http.StatusUnauthorized,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}

		for key, value := range tokenClaims {
			if key == "admin" {
				isAdmin = value.(bool)
			}
		}

		if isAdmin == false {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
