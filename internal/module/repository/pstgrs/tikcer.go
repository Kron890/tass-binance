package pstgrs

import (
	"fmt"
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

// Смотрит есть ли данные в бд (handler AddTicker)
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

// Добавляет данные в бд (handler AddTicker)
func (r *TickerRepository) AddTicker(ticker string, price float64) error {
	newTicker := models.TickerDb{
		Ticker: ticker,
		Price:  price,
	}

	if err := r.db.DB.Create(&newTicker).Error; err != nil {
		return err
	}
	return nil

}

// tickerDiff
func (r *TickerRepository) GetHistoryTikcer(ticker string, dateFrom int64, dateTo int64) (float64, float64, error) {
	// найти если есть данные в бд
	var startPrice models.TikcerHistory
	if err := r.db.DB.Where("ticker=? AND timestamp = ?", ticker, dateFrom).First(&startPrice).Error; err != nil {
		log.Println("нет даты в бд(data_from)")
	}
	var endPrice models.TikcerHistory
	if err := r.db.DB.Where("ticker = ? AND timestamp = ?", ticker, dateTo).First(&endPrice).Error; err != nil {
		log.Println("нет даты в бд(data_to)")
	}
	// вернуть
	return startPrice.Price, endPrice.Price, nil
}

// Получаем все тикеры в бд (UpdTicker)
func (r *TickerRepository) GetTicker() ([]models.TickerDb, error) { //изменил []string на []models.TickerDb
	var tickers []models.TickerDb
	return tickers, r.db.DB.Select("ticker").Find(&tickers).Error
}

// Обновляем тикеры полученные из бинанса в бд (UpdTicker)
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
				log.Println(err)
				return fmt.Errorf("can't create ticker history") //нет смысла выводить пользователю ошибку
			}

		}
	}
	return nil
}

// Вытаскивает из запроса название тикера (UpdTicker)
func BindTicker(c echo.Context) (string, error) {
	var ticker models.TickerDb
	if err := c.Bind(&ticker); err != nil {
		return "", err
	}
	return ticker.Ticker, nil
}
