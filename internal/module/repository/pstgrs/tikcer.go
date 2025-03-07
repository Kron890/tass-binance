package pstgrs

import "tass-binance/infrastructure/database"

type TickerRepository struct {
	db *database.DataBase
}

func NewRepo(db *database.DataBase) *TickerRepository {
	return &TickerRepository{db: db}
}
