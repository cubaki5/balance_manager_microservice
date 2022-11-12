package adapter

const (
	accrue            = "начисление"
	accrueComment     = "источник пополнения"
	acceptance        = "списание"
	acceptanceComment = "оплата заказа №%d по покупке %d"
	rejection         = "возврат"
	rejectionComment  = "возврат стоимости заказа №%d по покупке %d"

	mysql          = "mysql"
	dataSourceName = "avito:password@tcp(db:3306)/usersbalance?parseTime=true"
)
