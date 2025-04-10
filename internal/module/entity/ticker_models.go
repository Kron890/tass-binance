package entity

type TickerDifferenceEntity struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Difference string  `json:"difference"`
}

type TickerAddRequest struct {
	Name string
}

type Message struct {
	Message string `json:"message"`
}

type TickerDiffRequest struct {
	Name     string
	DateFrom string
	DateTo   string
}
