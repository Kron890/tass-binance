package module

type UseCase interface {
	ProcessTicker(ticker string) error
	TickerUpdateProcess() error
}
