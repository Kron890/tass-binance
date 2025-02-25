package repository

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (db *DataBase) connectBd() {
	dsn := "host=localhost user=user password=password dbname=mydatabase port=5432 sslmode=disable"
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error the database was not connected:%v", err)
	}
	db.DB = connect
	if err := db.DB.AutoMigrate(&TickerBd{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	if err := db.DB.AutoMigrate(&TikcerHistory{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
}
