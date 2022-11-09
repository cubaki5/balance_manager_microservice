package report_echo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"balance_avito/internal/models"
)

type Report interface {
	Accounting(ctx context.Context, reportDate models.ReportDate) (string, error)
	TransactionsHistory() error
}

type Handler struct {
	r Report
}

func NewReportHandler(r Report) *Handler {
	return &Handler{r}
}

func (h *Handler) Accounting(ctx echo.Context) error {
	var reportDate models.ReportDate

	err := ctx.Bind(&reportDate)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	report, err := h.r.Accounting(ctx.Request().Context(), reportDate)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusOK, report)
}

func (h *Handler) TransactionsHistory(ctx echo.Context) error {
	return nil
}
