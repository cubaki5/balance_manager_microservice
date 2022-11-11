package fund_echo

import (
	"context"
	"net/http"

	"balance_avito/internal/models"
	"balance_avito/internal/models/types"

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
	fund Fund
}

func NewFundHandler(f Fund) *Handler {
	return &Handler{f}
}

// @Summary     Accrue
// @Tags        fund
// @Description accrual income to the account with given user_id
// @Accept      json
// @Produce     json
// @Param       input body models.Accrual true "user_id and income to accrual"
// @Success     200
// @Failure     400 {object} error
// @Router      /fund/accrual [post]
func (h *Handler) Accrue(ctx echo.Context) error {
	var accrual models.Accrual

	err := ctx.Bind(&accrual)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.fund.Accrue(ctx.Request().Context(), accrual)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

// @Summary     Reservation
// @Tags        fund
// @Description reserve money on account with given user_id for given order_id and service_id
// @Accept      json
// @Produce     json
// @Param       input body models.Reservation true "user_id, service_id, order_id and cost to reserve"
// @Success     200
// @Failure     400 {object} error
// @Router      /fund/reservation [post]
func (h *Handler) Reservation(ctx echo.Context) error {
	var reservation models.Reservation

	err := ctx.Bind(&reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.fund.Reservation(ctx.Request().Context(), reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

// @Summary     AcceptPayment
// @Tags        fund
// @Description accept payment and save it in history
// @Accept      json
// @Produce     json
// @Param       input body models.Reservation true "user_id, service_id, order_id and cost to accept"
// @Success     200
// @Failure     400 {object} error
// @Router      /fund/payment_acceptance [post]
func (h *Handler) AcceptPayment(ctx echo.Context) error {
	var reservation models.Reservation

	err := ctx.Bind(&reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.fund.AcceptPayment(ctx.Request().Context(), reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

// @Summary     RejectPayment
// @Tags        fund
// @Description reject payment, returns funds to account and save it in history
// @Accept      json
// @Produce     json
// @Param       input body models.Reservation true "user_id, service_id, order_id and cost to reject"
// @Success     200
// @Failure     400 {object} error
// @Router      /fund/payment_rejection [post]
func (h *Handler) RejectPayment(ctx echo.Context) error {
	var reservation models.Reservation

	err := ctx.Bind(&reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.fund.RejectPayment(ctx.Request().Context(), reservation)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

// @Summary     GetBalance
// @Tags        fund
// @Description returns balance of account with given user_id
// @Accept      json
// @Produce     json
// @Param       input body models.Account true "user_id to get balance"
// @Success     200 {integer} integer
// @Failure     400 {object} error
// @Router      /fund/balance [get]
func (h *Handler) GetBalance(ctx echo.Context) error {
	var account models.Account

	err := ctx.Bind(&account)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	account.Balance, err = h.fund.GetBalance(ctx.Request().Context(), account)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]int32{"balance": account.Balance.Int32()})

}
