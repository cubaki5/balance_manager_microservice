package fund_handlers

import (
	"context"

	"balance_avito/internal/models"
)

type Database interface {
	Accrue() error
	Reservation() error
	AcceptPayment() error
	RejectPayment() error
	GetBalance() error
}

type Fund struct {
	db Database
}

func NewFund(db Database) *Fund {
	return &Fund{db}
}

func (f *Fund) Accrue(ctx context.Context, accrual models.Accrual) error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) Reservation(ctx context.Context, reservation models.Reservation) error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) AcceptPayment(ctx context.Context, reservation models.Reservation) error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) RejectPayment(ctx context.Context, reservation models.Reservation) error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) GetBalance(ctx context.Context, account models.Account) (models.Account, error) {
	//TODO implement me
	panic("implement me")
}
