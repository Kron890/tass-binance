package repository

import (
	"gorm.io/gorm"
	"main.go/cmd/internal/domain/models"
)

// Интерфейс репозитория для тикеров
type TickerRepository interface {
	GetTicker(name string) (*models.TickerModel, error)
	SaveTicker(ticker *models.TickerModel) error
}

// Реализация репозитория
type tickerRepo struct {
	db *gorm.DB
}

func NewTickerRepository(db *gorm.DB) TickerRepository {
	return &tickerRepo{db: db}
}

func (r *tickerRepo) GetTicker(name string) (int64, error) {
	var ticker models.TickerModel
	err := r.db.Model(tickerDB{}).Where("ticker=?", NameTicker).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return &ticker, err
}

//переделать GetTicker для проверки count
//сделать функцию для data

func (r *tickerRepo) SaveTicker(ticker *models.TickerModel) error {
	return r.db.Create(ticker).Error
}

//сделать функцию для добавления в бд с историей
//сделатьй функию для проверки в бд (count)
