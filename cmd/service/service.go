package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"balance_avito/internal/handlers/fund_handlers"
	"balance_avito/internal/handlers/report_handlers"
	"balance_avito/internal/infrastructure/echo_framework/fund_echo"
	"balance_avito/internal/infrastructure/echo_framework/report_echo"
	gocsv_adapter "balance_avito/internal/infrastructure/gocsv/adapter"
	"balance_avito/internal/infrastructure/sql_database/adapter"
)

func main() {
	db, err := adapter.RunDB()
	if err != nil {
		log.Fatal(err)
	}

	dbAdapter := adapter.NewDatabaseAdapter(db)
	goCSVAdapter := gocsv_adapter.NewGoSCVAdapter()

	fund := fund_handlers.NewFund(dbAdapter)
	report := report_handlers.NewReport(dbAdapter, goCSVAdapter)

	fundHandler := fund_echo.NewFundHandler(fund)
	reportHandler := report_echo.NewReportHandler(report)

	e := echo.New()

	fundRouter := e.Group("/fund")
	fundRouter.POST("/accrual", fundHandler.Accrue)
	fundRouter.POST("/reservation", fundHandler.Reservation)
	fundRouter.POST("/payment_acceptance", fundHandler.AcceptPayment)
	fundRouter.POST("/payment_rejection", fundHandler.RejectPayment)
	fundRouter.POST("/balance", fundHandler.GetBalance)

	reportRouter := e.Group("/report")
	reportRouter.POST("/accounting", reportHandler.Accounting)
	reportRouter.POST("/transactions_history", reportHandler.TransactionsHistory)

	e.Logger.Fatal(e.Start(":1323"))
}
