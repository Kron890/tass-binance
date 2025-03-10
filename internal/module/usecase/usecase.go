package usecase

import (
	"fmt"
	"tass-binance/internal/module"
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

func (u *UseCase) ProcessTickerDiff(ticker string, time_from string, time_do string) (float64, float64, error) {
	// преобразуем время - отдельная функция

	//ищем данные по дате - отдельная функция (в бд)

	//высчитываем diff -отдельная функция

	//возварщаем

	return 0, 0, nil
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
