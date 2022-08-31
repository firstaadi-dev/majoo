package main

import (
	"github.com/firstaadi-dev/majoo-backend-test/auth/delivery/http"
	"github.com/firstaadi-dev/majoo-backend-test/auth/repository/mysql"
	"github.com/firstaadi-dev/majoo-backend-test/auth/usecase"
	"github.com/firstaadi-dev/majoo-backend-test/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := config.DatabaseConnect()
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	authRepo := mysql.NewMysqlUserRepository(db)
	authUcase := usecase.NewAuthUsecase(
		authRepo,
		"saltpassword",
		[]byte("signingkey"),
		3600,
	)

	http.NewAuthHandler(e, authUcase)

	e.Start(":8080")
}
