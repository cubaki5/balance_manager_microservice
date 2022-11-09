package report_handlers

import (
	"context"

	"balance_avito/internal/models"
)

type GoCSVAdapter interface {
	MarshalStructToCSV([]models.AccountingReport) (string, error)
}

type Database interface {
	Accounting(ctx context.Context, reportDate models.ReportDate) ([]models.AccountingReport, error)
	TransactionsHistory() error
}

type Report struct {
	db    Database
	gocsv GoCSVAdapter
}

func NewReport(db Database, gocsv GoCSVAdapter) *Report {
	return &Report{db, gocsv}
}

func (r *Report) Accounting(ctx context.Context, reportDate models.ReportDate) (string, error) {
	accountReports, err := r.db.Accounting(ctx, reportDate)
	if err != nil {
		return "", err
	}

	report, err := r.gocsv.MarshalStructToCSV(accountReports)
	if err != nil {
		return "", err
	}

	return report, nil
}

func (r *Report) TransactionsHistory() error {
	//TODO implement me
	panic("implement me")
}
