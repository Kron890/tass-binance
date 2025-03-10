package pstgrs

import (
	"log"
	"tass-binance/infrastructure/database"
	"tass-binance/internal/module/models"

	"github.com/labstack/echo/v4"
)

type TickerRepository struct {
	db *database.DataBase
}

func NewRepo(db *database.DataBase) *TickerRepository {
	return &TickerRepository{db: db}
}

func (r *TickerRepository) CheckTicker(ticker string) (bool, error) {
	var count int64
	if err := r.db.DB.Model(&models.TickerDb{}).Where("ticker=?", ticker).Count(&count).Error; err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (r *TickerRepository) AddTickerDataBase(ticker string, price float64) error {
	newTicker := models.TickerDb{
		Ticker: ticker,
		Price:  price,
	}

	if err := r.db.DB.Create(&newTicker).Error; err != nil {
		return err
	}
	return nil

}

func BindTicker(c echo.Context) (string, error) {
	var ticker models.TickerDb
	if err := c.Bind(&ticker); err != nil {
		return "", err
	}
	return ticker.Ticker, nil
}

func (r *TickerRepository) GetTicker() ([]string, error) {
	var tickers []models.TickerDb
	var result []string
	if err := r.db.DB.Select("ticker").Find(&tickers).Error; err != nil {
		return nil, err
	}
	for _, t := range tickers {
		result = append(result, t.Ticker)
	}
	return result, nil
}

func (r *TickerRepository) UpdateTickerDb(tickers map[string]map[float64]int64) error {

	for ticker, data := range tickers {
		for price, timestamp := range data {
			if err := r.db.DB.Model(&models.TickerDb{}).Where("ticker = ?", ticker).Update("price", price).Error; err != nil {
				log.Println("error in data branching")
				return err
			}

			tickerHistory := models.TikcerHistory{
				Ticker:    ticker,
				Price:     price,
				Timestamp: timestamp}

			if err := r.db.DB.Create(&tickerHistory).Error; err != nil {
				log.Println("Error in data creation")
				return err
			}

		}
	}
	return nil
}
