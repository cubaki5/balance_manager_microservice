package main

import (
	"github.com/labstack/echo/v4"

	"balance_avito/handlers/fund_handlers"
	"balance_avito/handlers/report_handlers"
	"balance_avito/infrastructure/echo_framework/fund_echo"
	"balance_avito/infrastructure/echo_framework/report_echo"
	"balance_avito/infrastructure/sql_database/adapter"
)

func main() {
	db := adapter.NewDatabase()

	fund := fund_handlers.NewFund(db)
	report := report_handlers.NewReport(db)

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
	reportRouter.POST("/transaction_history", reportHandler.TransactionHistory)

	e.Logger.Fatal(e.Start(":1323"))
}
