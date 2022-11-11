package models

import (
	"errors"
	"time"

	"balance_avito/internal/models/types"
)

type Accrual struct {
	UserID types.UserID `json:"user_id"`
	Income types.Income `json:"income"`
}

type Reservation struct {
	UserID    types.UserID  `json:"user_id"`
	ServiceID types.UserID  `json:"service_id"`
	OrderID   types.OrderID `json:"order_id"`
	Cost      types.Cost    `json:"cost"`
}

type Account struct {
	UserID  types.UserID `json:"user_id"`
	Balance types.Balance
}

type ReportDate struct {
	Year  types.Year  `json:"year"`
	Month types.Month `json:"month"`
}

type AccountingReport struct {
	ServiceID types.ServiceID
	Income    types.Income
}

type TransactionsHistoryParams struct {
	UserID   types.UserID `json:"user_id"`
	SortDate int
	SortSum  int
	Page     int
	Limit    int
}

type TransactionsHistory struct {
	Operation types.Operation
	Comments  types.Comment
	Time      time.Time
	Sum       types.Income
}

func (r *Reservation) Validate() error {
	if r.UserID.Int64() <= 0 {
		return errors.New("неверное значение user_id")
	}

	if r.OrderID.Int64() <= 0 {
		return errors.New("неверное значение order_id")
	}

	if r.ServiceID.Int64() <= 0 {
		return errors.New("неверное значение service_id")
	}

	if r.Cost.Int32() <= 0 {
		return errors.New("неверное значение cost")
	}

	return nil
}

func (a *Accrual) Validate() error {
	if a.UserID.Int64() <= 0 {
		return errors.New("неверное значение user_id")
	}
	if a.Income.Int32() <= 0 {
		return errors.New("неверное значение income")
	}

	return nil
}

func (a Account) Validate() error {
	if a.UserID.Int64() <= 0 {
		return errors.New("неверное значение user_id")
	}
	return nil
}

func (t *TransactionsHistoryParams) Validate() error {
	if t.UserID.Int64() <= 0 {
		return errors.New("неверное значение user_id")
	}
	if t.Limit <= 0 {
		return errors.New("неверное значение limit")
	}
	if t.SortDate != 0 && t.SortDate != 1 {
		return errors.New("неверное значение sortDate (sortDate=1 - включение сортировки по дате, sortDate=0 - выключение сортировки по дате)")
	}
	if t.SortSum != 0 && t.SortSum != 1 {
		return errors.New("неверное значение sortSum (sortSum=1 - включение сортировки по сумме, sortSum=0 - выключение сортировки по сумме)")
	}
	if t.Page <= 0 {
		return errors.New("неверное значение page")
	}
	return nil
}

func (r *ReportDate) Validate() error {
	if r.Month > 12 || r.Month < 1 {
		return errors.New("неверное значение month")
	}
	if r.Year < 1900 || r.Year > 3000 {
		return errors.New("неверное значение year")
	}
	return nil
}
