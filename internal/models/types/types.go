package types

import "strconv"

type (
	UserID    int64
	Income    int32
	ServiceID int64
	OrderID   int64
	Cost      int32
	AccountID int64
	Balance   int32
	Year      int32
	Month     int32
)

func (i Income) String() string {
	return strconv.Itoa(int(i))
}

func (i ServiceID) String() string {
	return strconv.Itoa(int(i))
}

func (i Month) Int32() int32 {
	return int32(i)
}

func (i Year) Int32() int32 {
	return int32(i)
}

func (i Balance) Int32() int32 {
	return int32(i)
}

func (i AccountID) Int64() int64 {
	return int64(i)
}

func (i Income) Int32() int32 {
	return int32(i)
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

func (i Cost) Int32() int32 {
	return int32(i)
}
