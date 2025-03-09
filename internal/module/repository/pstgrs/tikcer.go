package pstgrs

import (
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

func (r TickerRepository) AddTickerDataBase(ticker string, price float64) error {
	newTicker := models.TickerDb{
		Ticker: ticker,
		Price:  price,
	}

	if err := r.db.DB.Create(&newTicker).Error; err != nil {
		return err
	}
	return nil

	//добавить вторую бд
}

func BindTicker(c echo.Context) (string, error) {
	var ticker models.TickerDb
	if err := c.Bind(&ticker); err != nil {
		return "", err
	}
	return ticker.Ticker, nil
}
