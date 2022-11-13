package fund_echo

type dtoAccrual struct {
	UserID int64 `json:"user_id"`
	Income int32 `json:"income"`
}

type dtoReservation struct {
	UserID    int64 `json:"user_id"`
	ServiceID int64 `json:"service_id"`
	OrderID   int64 `json:"order_id"`
	Cost      int32 `json:"cost"`
}

type dtoAccount struct {
	UserID int64 `json:"user_id"`
}
