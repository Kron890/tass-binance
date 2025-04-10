package module

import "tass-binance/internal/module/entity"

type UseCase interface {
	AddTicker(ticker entity.TickerAddRequest) error
	ProcessTickerDiff(ticker entity.TickerDiffRequest) (*entity.TickerDifferenceEntity, error)
}
