package models

type TickerDb struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price" `
}

type TikcerHistory struct {
	Ticker    string  `json:"ticker"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}
