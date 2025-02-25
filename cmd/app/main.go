package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/fetch/:ticker/:date_from/:date_to", GetFetch)
	e.POST("/add_ticker", PostAddTicker)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
