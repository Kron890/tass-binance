package extrenal_api

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
)

type Api struct {
	client *binance.Client
}

// type RegularApi struct {
// 	client *binance.Client
// }

func NewApi() *Api {
	client := binance.NewClient("", "")
	return &Api{client: client}
}

// func RegularNewApi() *RegularApi {
// 	client := binance.NewClient("", "")
// 	return &RegularApi{client: client}
// }

// получение прайса
func (a *Api) GetPrice(symbol string) (float64, error) {
	prices, err := a.client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return 0, err
	}
	price, err := strconv.ParseFloat(prices[0].Price, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}

// берем данные из бинанса
func (a *Api) GetRegularPrice(tickers []string) (map[string]map[float64]int64, error) {
	tickersBinance := make(map[string]map[float64]int64, len(tickers))
	for _, ticker := range tickers {
		priceStr, err := a.client.NewListPricesService().Symbol(ticker).Do(context.Background())
		if err != nil {
			return nil, fmt.Errorf("error getting price for %s: %w", ticker, err)

		}

		if len(priceStr) == 0 {
			return nil, fmt.Errorf("no price data for %s", ticker)
		}

		price, err := strconv.ParseFloat(priceStr[0].Price, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing price for %s: %w", ticker, err)
		}
		tickersBinance[ticker] = map[float64]int64{
			price: time.Now().Unix(),
		}
	}

	return tickersBinance, nil

}
