package report_echo

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"balance_avito/internal/models"
)

type Report interface {
	Accounting(ctx context.Context, reportDate models.ReportDate) (string, error)
	TransactionsHistory(ctx context.Context, transactionsHistoryParams models.TransactionsHistoryParams) (string, error)
}

type Handler struct {
	r Report
}

func NewReportHandler(r Report) *Handler {
	return &Handler{r}
}

// @Summary     Accounting
// @Tags        report
// @Description returns an accounting report of giving date
// @Accept      json
// @Produce     json
// @Param       input body models.ReportDate true "specific year and month"
// @Success     200 {string} string
// @Failure     400 {object} error
// @Router      /report/accounting [get]
func (h *Handler) Accounting(ctx echo.Context) error {
	var reportDate models.ReportDate

	err := ctx.Bind(&reportDate)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	report, err := h.r.Accounting(ctx.Request().Context(), reportDate)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusOK, report)
}

// @Summary     TransactionsHistory
// @Tags        report
// @Description returns a transaction history of account with given user_id
// @Accept      json
// @Produce     json
// @Param		input query models.HistoryQueryParams true "Sorts, page and limit"
// @Param       input body models.HistoryBodyParams true "user_id to get history"
// @Success     200 {string} string
// @Failure     400 {object} error
// @Router      /report/transactions_history [get]
func (h *Handler) TransactionsHistory(ctx echo.Context) error {
	queryParamsMap := ctx.QueryParams()

	queryParams, err := marshalQueryParamsInStruct(queryParamsMap)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = ctx.Bind(&queryParams)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	transactionsHistory, err := h.r.TransactionsHistory(ctx.Request().Context(), queryParams)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusOK, transactionsHistory)
}

func marshalQueryParamsInStruct(queryParamsMap url.Values) (models.TransactionsHistoryParams, error) {
	var queryParams = models.TransactionsHistoryParams{}

	for key, queryParam := range queryParamsMap {
		intQueryParam, err := strconv.Atoi(queryParam[0])
		if err != nil {
			return models.TransactionsHistoryParams{}, err
		}

		switch key {
		case "sortDate":
			queryParams.SortDate = intQueryParam
		case "sortSum":
			queryParams.SortSum = intQueryParam
		case "limit":
			queryParams.Limit = intQueryParam
		case "page":
			queryParams.Page = intQueryParam
		}
	}
	return queryParams, nil
}
