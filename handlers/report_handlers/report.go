package report_handlers

type Database interface {
	Accounting() error
	TransactionHistory() error
}

type Report struct {
	db Database
}

func NewReport(db Database) *Report {
	return &Report{db}
}

func (r *Report) Accounting() error {
	//TODO implement me
	panic("implement me")
}

func (r *Report) TransactionHistory() error {
	//TODO implement me
	panic("implement me")
}
