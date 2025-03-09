package database

import (
	"log"
	"os"
	"tass-binance/internal/module/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	DB *gorm.DB
}

func NewDbConnection() (*DataBase, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=user password=password dbname=mydatabase port=5432 sslmode=disable"
	}
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := connect.AutoMigrate(&models.TickerDb{}, &models.TikcerHistory{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	return &DataBase{DB: connect}, nil
}
