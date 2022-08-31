package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/firstaadi-dev/majoo-backend-test/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	useCase auth.UseCase
}

func NewAuthHandler(r *echo.Echo, us auth.UseCase) {
	handler := &AuthHandler{
		useCase: us,
	}
	r.POST("/login", handler.Login)
}

type loginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var input = new(loginInput)

	err := c.Bind(&input)
	if err != nil {
		fmt.Println(err)
		return err
	}
	token, err := h.useCase.Login(input.Username, input.Password)
	if err == sql.ErrNoRows {
		return c.String(http.StatusUnauthorized, "invalid username or password")
	}
	return c.JSON(http.StatusOK, loginResponse{
		Token: token,
	})
}
