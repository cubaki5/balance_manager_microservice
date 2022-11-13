package report_echo

type dtoReportDate struct {
	Year  int32 `json:"year"`
	Month int32 `json:"month"`
}

type DtoHistoryQueryParams struct {
	SortDate int
	SortSum  int
	Page     int
	Limit    int
}

type dtoHistoryBodyParams struct {
	UserID int64 `json:"user_id"`
}
