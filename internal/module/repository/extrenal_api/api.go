package extrenal_api

import (
	"context"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

type Api struct {
	client *binance.Client
}

func NewApi() *Api {
	client := binance.NewClient("", "")
	return &Api{client: client}
}

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
