package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var db DataBase
	db.connectBd()
	go db.RegularUpd()
	e.GET("/fetch/:ticker/:date_from/:date_to", db.GetFetch)
	e.POST("/add_ticker", db.PostAddTicker)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
