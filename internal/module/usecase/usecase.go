package usecase

import (
	"fmt"
	"tass-binance/internal/module"
	"tass-binance/internal/module/usecase/helpers"
)

type UseCase struct {
	dbRepo          module.ModuleDatabase
	externalApiRepo module.ModuleExternalService
}

func NewUseCase(dbRepo module.ModuleDatabase, externalApiRepo module.ModuleExternalService) *UseCase {

	return &UseCase{
		dbRepo:          dbRepo,
		externalApiRepo: externalApiRepo}
}

func (u *UseCase) ProcessTicker(ticker string) error {
	exists, err := u.dbRepo.CheckTicker(ticker)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the ticker is already in the database")
	}

	price, err := u.externalApiRepo.GetPrice(ticker)
	if err != nil {
		return err
	}
	err = u.dbRepo.AddTickerDataBase(ticker, price)

	if err != nil {
		return err
	}
	return nil

}

// GetTickerDiff
func (u *UseCase) ProcessTickerDiff(ticker string, dateFrom string, dateTo string) (float64, float64, error) {

	timeFrom, timeTo, err := helpers.ConvertTime(dateFrom, dateTo)
	if err != nil {
		return 0, 0, err
	}

	startPrice, endPrice, _ := u.dbRepo.GetHistoryTikcer(ticker, timeFrom, timeTo)

	percDiff := helpers.DiffCalculator(endPrice, startPrice)

	return percDiff, endPrice, nil
}

func (u *UseCase) TickerUpdateProcess() error {

	tickers, err := u.dbRepo.GetTicker()
	if err != nil {
		return err
	}

	tickersBinance, err := u.externalApiRepo.GetRegularPrice(tickers)
	if err != nil {
		return err
	}

	err = u.dbRepo.UpdateTickerDb(tickersBinance)
	if err != nil {
		return err
	}
	return nil
}
