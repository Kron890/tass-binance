package usecase

import (
	"fmt"
	"tass-binance/internal/module"
	"tass-binance/internal/module/entity"
	"tass-binance/internal/module/models"
	"tass-binance/internal/module/usecase/helpers"
	"tass-binance/pkg/logger"
	"tass-binance/pkg/time_convert"
	"time"

	"github.com/labstack/gommon/log"
)

type UseCase struct {
	dbRepo          module.ModuleDatabase
	externalApiRepo module.ModuleExternalService
}

func NewUseCase(dbRepo module.ModuleDatabase, externalApiRepo module.ModuleExternalService) *UseCase {

	uc := &UseCase{
		dbRepo:          dbRepo,
		externalApiRepo: externalApiRepo}
	go uc.tickerUpdateRoutine()
	return uc
}

// processTicker
func (u *UseCase) AddTicker(ticker entity.TickerAddRequest) error {
	l := logger.NewLogger()

	exists, err := u.dbRepo.CheckTicker(ticker.Name)
	if err != nil {
		l.Errorf("CheckTicker: %s", err)
		return err
	}
	if exists {
		l.Errorf("CheckTicke: the ticker is already in the database")
		return fmt.Errorf("the ticker is already in the database")
	}

	price, err := u.externalApiRepo.GetPrice(ticker.Name)
	if err != nil {
		l.Infof("GetPricee: %s", err)
		return err
	}
	return u.dbRepo.AddTicker(ticker.Name, price)

}

// GetTickerDiff
func (u *UseCase) ProcessTickerDiff(ticker entity.TickerDiffRequest) (*entity.TickerDifferenceEntity, error) {

	timeFrom, timeTo, err := time_convert.ConvertTime(ticker.DateFrom, ticker.DateTo)
	if err != nil {
		return &entity.TickerDifferenceEntity{}, err
	}

	startPrice, endPrice, _ := u.dbRepo.GetHistoryTikcer(ticker.Name, timeFrom, timeTo)

	percDiff := helpers.DiffCalculator(endPrice, startPrice)

	difference := &entity.TickerDifferenceEntity{
		Name:       ticker.Name,
		Price:      endPrice,
		Difference: fmt.Sprintf("%.3f%%", percDiff),
	}
	return difference, nil
}

func (u *UseCase) tickerUpdateRoutine() error {
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-tick.C:
			err := u.tickerUpdateProcess()
			if err != nil {
				log.Warnf("ticker update error: %s", err)
			}
		}
	}
}

func (u *UseCase) tickerUpdateProcess() error {
	var ticker []models.TickerDb
	ticker, err := u.dbRepo.GetTicker()
	if err != nil {
		return err
	}
	tickers := make([]string, 0, len(ticker))
	for _, name := range ticker {
		tickers = append(tickers, name.Ticker)
	}

	tickersBinance, err := u.externalApiRepo.GetRegularPrice(tickers)
	if err != nil {
		return err
	}

	return u.dbRepo.UpdateTickerDb(tickersBinance)

}
