package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetIDFromContext(c echo.Context) (int, error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user := claims["user"].(map[string]interface{})
	id := user["ID"].(float64)
	if id == 0 {
		return 0, fmt.Errorf("invalid user id")
	}
	return int(id), nil
}
