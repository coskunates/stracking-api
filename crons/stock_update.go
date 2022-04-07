package crons

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"stock/services"
	"stock/utils/logger_utils"
	"strconv"
	"time"
)

const (
	fcsApiAccessToken = "fcs_api_access_token"
	stockPriceUrl     = "https://fcsapi.com/api-v3/stock/latest?id=%d&access_key=%s"
)

var (
	myClient    = &http.Client{Timeout: 10 * time.Second}
	accessToken = os.Getenv(fcsApiAccessToken)
)

type FcsApiStockPriceResponse struct {
	Status   bool                           `json:"status"`
	Code     int64                          `json:"code"`
	Msg      string                         `json:"msg"`
	Response []FcsApiStockPriceResponseData `json:"response"`
}

type FcsApiStockPriceResponseData struct {
	CurrentPrice string `json:"c"`
	High         string `json:"h"`
	Low          string `json:"l"`
}

func SetStockCurrentPrice() {
	summaryService := services.NewSummaryService()
	stockService := services.NewStockService()

	// Only updates open positions' stock prices
	portfolio, _ := summaryService.Portfolio()

	response := FcsApiStockPriceResponse{}
	for _, position := range *portfolio {
		_ = getPrice(position.Stock.FcsApiId, &response)

		currentPrice, _ := strconv.ParseFloat(response.Response[0].CurrentPrice, 32)
		stockService.UpdateStockCurrentPrice(position.Stock.ID, currentPrice)
	}
}

func getPrice(fcsApiId uint64, target interface{}) error {
	resp, err := myClient.Get(fmt.Sprintf(stockPriceUrl, fcsApiId, accessToken))
	if err != nil {
		log.Fatalln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger_utils.Error("Body close defer error", err)
		}
	}(resp.Body)

	return json.NewDecoder(resp.Body).Decode(target)
}
