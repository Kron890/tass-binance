package database

import (
	"fmt"
	"log"
	"tass-binance/config"
	"tass-binance/internal/module/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	DB *gorm.DB
}

func NewDbConnection(cfg config.Config) (*DataBase, error) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.PostgresConfig.User, cfg.PostgresConfig.Password, cfg.PostgresConfig.NameDb, cfg.PostgresConfig.Port)
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := connect.AutoMigrate(&models.TickerDb{}, &models.TikcerHistory{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	return &DataBase{DB: connect}, nil
}
