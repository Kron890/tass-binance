package infrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/cmd/internal/domain/models"
)

type DataBase struct {
	DB *gorm.DB
}

func (c *DataBase) Connect() error {
	dsn := "host=localhost user=user password=password dbname=mydatabase port=5432 sslmode=disable"
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error the database was not connected:%v", err)
	}
	c.DB = connect
	if err := c.DB.AutoMigrate(&models.TickerModel{}); err != nil {
		return fmt.Errorf("migration error: %v", err)

	}
	if err := c.DB.AutoMigrate(&models.TickerPriceRecord{}); err != nil {
		return fmt.Errorf("migration error: %v", err)

	}
	return nil
}
