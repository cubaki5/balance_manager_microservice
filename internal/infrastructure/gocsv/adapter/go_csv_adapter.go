package adapter

import (
	"encoding/csv"
	"io"

	"balance_avito/internal/models"

	"github.com/gocarina/gocsv"
)

type GoSCVAdapter struct {
}

func NewGoSCVAdapter() *GoSCVAdapter {
	return &GoSCVAdapter{}
}

func (g *GoSCVAdapter) MarshalStructReportToCSV(accountingReport []models.AccountingReport) (string, error) {
	report, err := marshalStructToCSV[models.AccountingReport](accountingReport)
	return report, err
}

func (g *GoSCVAdapter) MarshalStructHistoryToCSV(transactionsHistory []models.TransactionsHistory) (string, error) {
	report, err := marshalStructToCSV[models.TransactionsHistory](transactionsHistory)
	return report, err
}

type parsedStruct interface {
	models.AccountingReport | models.TransactionsHistory
}

func marshalStructToCSV[T parsedStruct](dbReport []T) (string, error) {
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = ';'
		return gocsv.NewSafeCSVWriter(writer)
	})

	report, err := gocsv.MarshalStringWithoutHeaders(&dbReport)
	if err != nil {
		return "", err
	}

	return report, err
}
