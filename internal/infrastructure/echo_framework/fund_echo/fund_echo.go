package fund_echo

import (
	"context"
	"net/http"

	"balance_avito/internal/models/types"

	"balance_avito/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Fund interface {
	Accrue(ctx context.Context, accrual models.Accrual) error
	Reservation(ctx context.Context, reservation models.Reservation) error
	AcceptPayment(ctx context.Context, reservation models.Reservation) error
	RejectPayment(ctx context.Context, reservation models.Reservation) error
	GetBalance(ctx context.Context, account models.Account) (types.Balance, error)
}

type Handler struct {
	f Fund
}

func NewFundHandler(f Fund) *Handler {
	return &Handler{f}
}

func (h *Handler) Accrue(ctx echo.Context) error {
	var accrual models.Accrual

	err := ctx.Bind(&accrual)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.f.Accrue(ctx.Request().Context(), accrual)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) Reservation(ctx echo.Context) error {
	var reservation models.Reservation

	err := ctx.Bind(&reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.f.Reservation(ctx.Request().Context(), reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) AcceptPayment(ctx echo.Context) error {
	var reservation models.Reservation

	err := ctx.Bind(&reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.f.AcceptPayment(ctx.Request().Context(), reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) RejectPayment(ctx echo.Context) error {
	var reservation models.Reservation

	err := ctx.Bind(&reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.f.RejectPayment(ctx.Request().Context(), reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) GetBalance(ctx echo.Context) error {
	var account models.Account

	err := ctx.Bind(&account)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	account.Balance, err = h.f.GetBalance(ctx.Request().Context(), account)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]int32{"balance": account.Balance.Int32()})

}
