package models

import (
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
	ID      types.AccountID
	UserID  types.UserID `json:"user_id"`
	Balance types.Balance
}
