package main

import (
	"fmt"

	"github.com/firstaadi-dev/majoo-backend-test/auth/delivery/http"
	"github.com/firstaadi-dev/majoo-backend-test/auth/repository/mysql"
	"github.com/firstaadi-dev/majoo-backend-test/auth/usecase"

	_transactionHandler "github.com/firstaadi-dev/majoo-backend-test/transaction/delivery/http"
	_transactionRepo "github.com/firstaadi-dev/majoo-backend-test/transaction/repository/mysql"
	_transactionUcase "github.com/firstaadi-dev/majoo-backend-test/transaction/usecase"

	"github.com/firstaadi-dev/majoo-backend-test/config"
	"github.com/firstaadi-dev/majoo-backend-test/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := config.DatabaseConnect()
	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("signingkey"),
	// }))

	e.GET("/ping", func(c echo.Context) error {
		id, err := helper.GetIDFromContext(c)
		if err != nil {
			return err
		}
		fmt.Println(id)
		return c.String(200, "ok")
	})

	loginGroup := e.Group("")
	authRepo := mysql.NewMysqlUserRepository(db)
	authUcase := usecase.NewAuthUsecase(
		authRepo,
		[]byte("signingkey"),
		3600,
	)

	reportGroup := e.Group("/report")
	transactionRepo := _transactionRepo.NewMysqlTransactionRepository(db)
	transactionUcase := _transactionUcase.NewTransactionUsecase(transactionRepo)

	http.NewAuthHandler(loginGroup, authUcase)
	_transactionHandler.NewTransactionHandler(reportGroup, transactionUcase)
	e.Start(":8080")
}
