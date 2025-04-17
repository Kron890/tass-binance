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
)

type UseCase struct {
	dbRepo          module.ModuleDatabase
	externalApiRepo module.ModuleExternalService
	l               *logger.Logger
}

func NewUseCase(dbRepo module.ModuleDatabase, externalApiRepo module.ModuleExternalService, l *logger.Logger) *UseCase {

	uc := &UseCase{
		dbRepo:          dbRepo,
		externalApiRepo: externalApiRepo,
		l:               l}
	go uc.tickerUpdateRoutine()
	return uc
}

// processTicker
func (u *UseCase) AddTicker(ticker entity.TickerAddRequest) error {
	exists, err := u.dbRepo.CheckTicker(ticker.Name)
	if err != nil {
		u.l.Errorf("CheckTicker: %s", err)
		return err
	}
	if exists {
		u.l.Errorf("CheckTicke: the ticker is already in the database")
		return fmt.Errorf("the ticker is already in the database")
	}

	price, err := u.externalApiRepo.GetPrice(ticker.Name)
	if err != nil {
		u.l.Infof("GetPricee: %s", err)
		return err
	}
	return u.dbRepo.AddTicker(ticker.Name, price)

}

// GetTickerDiff
func (u *UseCase) ProcessTickerDiff(ticker entity.TickerDiffRequest) (*entity.TickerDifferenceEntity, error) {

	timeFrom, timeTo, err := time_convert.ConvertTime(ticker.DateFrom, ticker.DateTo)
	if err != nil {
		u.l.Errorf("convertTime: %s", err.Error())
		return &entity.TickerDifferenceEntity{}, err
	}

	startPrice, endPrice, _ := u.dbRepo.GetHistoryTikcer(ticker.Name, timeFrom, timeTo)

	percDiff := helpers.DiffCalculator(endPrice, startPrice)

	difference := &entity.TickerDifferenceEntity{
		Name:       ticker.Name,
		Price:      endPrice,
		Difference: fmt.Sprintf("%.3f%%", percDiff),
	}
	u.l.Infof("processTikcerDiff: successfully")
	return difference, nil
}

func (u *UseCase) tickerUpdateRoutine() error {
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-tick.C:
			err := u.tickerUpdateProcess()
			if err != nil {
				u.l.WarnF("ticker update error: %s", err)
			}
		}
	}
}

func (u *UseCase) tickerUpdateProcess() error {
	var ticker []models.TickerDb
	ticker, err := u.dbRepo.GetTicker()
	if err != nil {
		u.l.Errorf("tickerUpdateProcess, %s", err.Error())
		return err
	}
	tickers := make([]string, 0, len(ticker))
	for _, name := range ticker {
		tickers = append(tickers, name.Ticker)
	}

	tickersBinance, err := u.externalApiRepo.GetRegularPrice(tickers)
	if err != nil {
		u.l.Errorf("getRegularPrice %s", err.Error())
		return err
	}

	return u.dbRepo.UpdateTickerDb(tickersBinance)

}
