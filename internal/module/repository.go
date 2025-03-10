package module

type ModuleDatabase interface {
	CheckTicker(ticker string) (bool, error)
	AddTickerDataBase(ticker string, price float64) error
	GetTicker() ([]string, error)
	UpdateTickerDb(map[string]map[float64]int64) error
}

type ModuleExternalService interface {
	GetPrice(ticker string) (float64, error)
	GetRegularPrice(tickers []string) (map[string]map[float64]int64, error)
}
