package report_echo

import "github.com/labstack/echo/v4"

type Report interface {
	Accounting() error
	TransactionHistory() error
}

type Handler struct {
	r Report
}

func NewReportHandler(r Report) *Handler {
	return &Handler{r}
}

func (h *Handler) Accounting(ctx echo.Context) error {
	return nil
}

func (h *Handler) TransactionHistory(ctx echo.Context) error {
	return nil
}
