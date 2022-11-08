package fund_handlers

import (
	"context"

	"balance_avito/internal/models/types"

	"balance_avito/internal/models"
)

type Database interface {
	Accrue(ctx context.Context, accrual models.Accrual) error
	Reservation(ctx context.Context, reservation models.Reservation) error
	AcceptPayment(ctx context.Context, reservation models.Reservation) error
	RejectPayment(ctx context.Context, reservation models.Reservation) error
	GetBalance(ctx context.Context, account models.Account) (models.Account, error)
}

type Fund struct {
	db Database
}

func NewFund(db Database) *Fund {
	return &Fund{db}
}

func (f *Fund) Accrue(ctx context.Context, accrual models.Accrual) error {
	return f.db.Accrue(ctx, accrual)
}

func (f *Fund) Reservation(ctx context.Context, reservation models.Reservation) error {
	return f.db.Reservation(ctx, reservation)
}

func (f *Fund) AcceptPayment(ctx context.Context, reservation models.Reservation) error {
	return f.db.AcceptPayment(ctx, reservation)
}

func (f *Fund) RejectPayment(ctx context.Context, reservation models.Reservation) error {
	return f.db.RejectPayment(ctx, reservation)
}

func (f *Fund) GetBalance(ctx context.Context, account models.Account) (types.Balance, error) {
	var err error
	account, err = f.db.GetBalance(ctx, account)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}
