package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"main.go/cmd/internal/infrastructure"
	"main.go/cmd/internal/repository"
	"main.go/cmd/internal/usecase"
)

func main() {
	e := echo.New()

	//исправить в будущем
	var db infrastructure.DataBase

	err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
		return
	}
	TickerRepo := repository.NewTickerService(db.DB)

	e.GET("/fetch/:ticker/:date_from/:date_to", usecase.GetFetch)
	e.POST("/add_ticker", usecase.PostAddTicker)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
