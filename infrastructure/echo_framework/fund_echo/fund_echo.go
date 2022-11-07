package fund_echo

import "github.com/labstack/echo/v4"

type Fund interface {
	Accrue() error
	Reservation() error
	AcceptPayment() error
	RejectPayment() error
	GetBalance() error
}

type Handler struct {
	f Fund
}

func NewFundHandler(f Fund) *Handler {
	return &Handler{f}
}

func (h *Handler) Accrue(ctx echo.Context) error {
	return nil
}

func (h *Handler) Reservation(ctx echo.Context) error {
	return nil
}

func (h *Handler) AcceptPayment(ctx echo.Context) error {
	return nil
}

func (h *Handler) RejectPayment(ctx echo.Context) error {
	return nil
}

func (h *Handler) GetBalance(ctx echo.Context) error {
	return nil
}
