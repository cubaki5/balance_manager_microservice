package models

import (
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
