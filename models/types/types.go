package types

type (
	UserID    int64
	Income    int64
	ServiceID int64
	OrderID   int64
	Cost      int64
	AccountID int64
	Balance   int64
)

func (i Balance) Int64() int64 {
	return int64(i)
}

func (i AccountID) Int64() int64 {
	return int64(i)
}

func (i Income) Int64() int64 {
	return int64(i)
}

func (i UserID) Int64() int64 {
	return int64(i)
}

func (i ServiceID) Int64() int64 {
	return int64(i)
}

func (i OrderID) Int64() int64 {
	return int64(i)
}

func (i Cost) Int64() int64 {
	return int64(i)
}
