package repository

import (
	"gorm.io/gorm"
	"main.go/cmd/internal/domain/models"
)

// Интерфейс репозитория для тикеров
type TickerRepository interface {
	CheckTicker(name string) (bool, error)
	SaveTicker(ticker *models.TickerModel) error
}

// Реализация репозитория
type tickerService struct {
	db *gorm.DB
}

func NewTickerService(db *gorm.DB) TickerRepository {
	return &tickerService{db: db}
}

func (r *tickerService) CheckTicker(name string) (bool, error) {
	var NameTicker models.TickerModel
	var count int64
	err := r.db.Model(tickerDB{}).Where("ticker=?", NameTicker).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

//переделать GetTicker для проверки count
//сделать функцию для data

func (r *tickerService) SaveTicker(ticker *models.TickerModel) error {
	return r.db.Create(ticker).Error
}

//сделать функцию для добавления в бд с историей
//сделатьй функию для проверки в бд (count)
