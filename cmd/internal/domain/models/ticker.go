package models

type Status struct {
	Message string `json:"message"`
}

type TickerModel struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price" `
}

type TickerPriceRecord struct {
	Ticker    string  `json:"ticker"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

type PriceDifference struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Difference string  `json:"difference"`
}
