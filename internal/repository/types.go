package repository

import "main.go/cmd/internal/domain/models"

//TickerModel
type tickerDB struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

func convTickerModel(t models.TickerModel) tickerDB {
	return tickerDB{
		Ticker: t.Ticker,
		Price:  t.Price,
	}
}

//TickerPriceRecord
type tickerHistory struct {
	Ticker    string  `json:"ticker"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

func convTickerPriceRecord(t models.TickerPriceRecord) tickerHistory {
	return tickerHistory{
		Ticker:    t.Ticker,
		Price:     t.Price,
		Timestamp: t.Timestamp,
	}
}

//PriceDifference
type priceDiff struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Difference string  `json:"difference"`
}

func convPriceDifference(p models.PriceDifference) priceDiff {
	return priceDiff{
		Name:       p.Name,
		Price:      p.Price,
		Difference: p.Difference,
	}
}
