package module

import "tass-binance/internal/module/models"

type ModuleDatabase interface {
	CheckTicker(ticker string) (bool, error)
	AddTicker(ticker string, price float64) error
	GetTicker() ([]models.TickerDb, error)
	UpdateTickerDb(map[string]map[float64]int64) error
	GetHistoryTikcer(ticker string, dateFrom int64, dateTo int64) (float64, float64, error)
}

type ModuleExternalService interface {
	GetPrice(ticker string) (float64, error)
	GetRegularPrice(tickers []string) (map[string]map[float64]int64, error)
}
