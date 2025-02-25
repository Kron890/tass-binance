package repository

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataBase struct {
	db *gorm.DB
}

func (c *dataBase) connectDB() {
	dsn := "host=localhost user=user password=password dbname=mydatabase port=5432 sslmode=disable"
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error the database was not connected:%v", err)
	}
	c.db = connect
	if err := c.db.AutoMigrate(&tickerDB{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	if err := c.db.AutoMigrate(&tickerHistory{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
}
