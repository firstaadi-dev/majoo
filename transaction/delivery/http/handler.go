package http

import (
	"net/http"
	"strconv"

	"github.com/firstaadi-dev/majoo-backend-test/helper"
	"github.com/firstaadi-dev/majoo-backend-test/transaction"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TransactionHandler struct {
	useCase transaction.UseCase
}

func NewTransactionHandler(r *echo.Group, us transaction.UseCase) {
	handler := &TransactionHandler{
		useCase: us,
	}
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("signingkey"),
	}))
	r.GET("/merchant/:id/:date", handler.GetMerchantReport)
	r.GET("/outlet/:id/:date", handler.GetOutletReport)
}

func (h *TransactionHandler) GetMerchantReport(c echo.Context) error {
	idToken, err := helper.GetIDFromContext(c)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if idToken != id {
		return c.String(http.StatusUnauthorized, "can't access another merchant report")
	}
	date, err := strconv.Atoi(c.Param("date"))
	if err != nil {
		return err
	}
	if date <= 0 || date > 30 {
		return c.String(http.StatusBadRequest, "invalid date")
	}
	res, err := h.useCase.ReportDailyMerchantOmzet(id, date)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)

}

func (h *TransactionHandler) GetOutletReport(c echo.Context) error {
	idToken, err := helper.GetIDFromContext(c)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	outletMerchantId, err := h.useCase.MerchantByOutletID(id)
	if err != nil {
		return c.String(http.StatusNotFound, "outlet not found")
	}
	if idToken != outletMerchantId.ID {
		return c.String(http.StatusUnauthorized, "can't access another merchant report")
	}
	date, err := strconv.Atoi(c.Param("date"))
	if err != nil {
		return err
	}
	if date <= 0 || date > 30 {
		return c.String(http.StatusBadRequest, "invalid date")
	}
	res, err := h.useCase.ReportDailyOutletOmzet(id, date)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
