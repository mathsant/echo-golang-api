package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"strings"
)

func GetTokensClaims(c echo.Context) (jwt.MapClaims, error) {
	tokenRequest := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(tokenRequest, "Bearer ")
	tokenRequest = splitToken[1]

	tokenDecoded, err := DecodeToken(tokenRequest)

	if err != nil {
		return nil, err
	}

	return tokenDecoded, nil
}
