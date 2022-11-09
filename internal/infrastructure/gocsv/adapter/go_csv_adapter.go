package adapter

import (
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"

	"balance_avito/internal/models"
)

type GoSCVAdapter struct {
}

func NewGoSCVAdapter() *GoSCVAdapter {
	return &GoSCVAdapter{}
}

func (g *GoSCVAdapter) MarshalStructToCSV(reports []models.AccountingReport) (string, error) {

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = ';'
		return gocsv.NewSafeCSVWriter(writer)
	})

	report, err := gocsv.MarshalStringWithoutHeaders(&reports)
	if err != nil {
		return "", err
	}

	return report, err
}
