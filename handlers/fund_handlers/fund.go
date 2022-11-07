package fund_handlers

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

func (f *Fund) Accrue() error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) Reservation() error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) AcceptPayment() error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) RejectPayment() error {
	//TODO implement me
	panic("implement me")
}

func (f *Fund) GetBalance() error {
	//TODO implement me
	panic("implement me")
}
