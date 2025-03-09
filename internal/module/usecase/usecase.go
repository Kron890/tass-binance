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
