package report_echo

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"balance_avito/internal/models/types"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"balance_avito/internal/models"
)

//go:generate mockgen --source=report_echo.go --destination=mocks/mock_report_echo.go --package=mocks

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
// @Param       input body dtoReportDate true "specific year and month"
// @Success     200 {string} string
// @Failure     400 {object} error
// @Router      /report/accounting [get]
func (h *Handler) Accounting(ctx echo.Context) error {
	var reportDateDTO dtoReportDate

	err := ctx.Bind(&reportDateDTO)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	reportDate := marshalReportDate(reportDateDTO)

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
// @Param		input query DtoHistoryQueryParams true "Sorts, page and limit"
// @Param       input body dtoHistoryBodyParams true "user_id to get history"
// @Success     200 {string} string
// @Failure     400 {object} error
// @Router      /report/transactions_history [get]
func (h *Handler) TransactionsHistory(ctx echo.Context) error {
	queryParamsDTO := ctx.QueryParams()

	var bodyParamsDTO dtoHistoryBodyParams

	err := ctx.Bind(&bodyParamsDTO)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	historyParams, err := marshalQueryAndBodyParams(queryParamsDTO, bodyParamsDTO)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	transactionsHistory, err := h.r.TransactionsHistory(ctx.Request().Context(), historyParams)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusOK, transactionsHistory)
}

func marshalQueryAndBodyParams(queryParamsMap url.Values, bodyParamsDTO dtoHistoryBodyParams) (models.TransactionsHistoryParams, error) {
	var historyParams = models.TransactionsHistoryParams{}

	for key, queryParam := range queryParamsMap {
		intQueryParam, err := strconv.Atoi(queryParam[0])
		if err != nil {
			return models.TransactionsHistoryParams{}, err
		}

		switch key {
		case "sortDate":
			historyParams.SortDate = intQueryParam
		case "sortSum":
			historyParams.SortSum = intQueryParam
		case "limit":
			historyParams.Limit = intQueryParam
		case "page":
			historyParams.Page = intQueryParam
		}
	}

	historyParams.UserID = types.UserID(bodyParamsDTO.UserID)
	return historyParams, nil
}

func marshalReportDate(reportDateDTO dtoReportDate) models.ReportDate {
	var reportDate models.ReportDate

	reportDate.Year = types.Year(reportDateDTO.Year)
	reportDate.Month = types.Month(reportDateDTO.Month)

	return reportDate
}
