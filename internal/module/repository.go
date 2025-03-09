package module

type ModuleDatabase interface {
	CheckTicker(ticker string) (bool, error)
	AddTickerDataBase(ticker string, price float64) error
}

type ModuleExternalService interface {
	GetPrice(ticker string) (float64, error)
}
