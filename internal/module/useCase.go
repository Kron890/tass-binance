package module

type UseCase interface {
	ProcessTicker(ticker string) error
	ProcessTickerDiff(ticker string, dateFrom string, dateTo string) (float64, float64, error)
	TickerUpdateProcess() error
}
