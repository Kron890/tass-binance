package deliv

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func MapRoutes(e *echo.Echo, h Handler) {

	e.POST("/add_ticker", h.AddTicker)
	e.GET("/fetch/:ticker/:date_from/:date_to", h.TickerDiff)
}

func StartUpd(e *echo.Echo, h Handler) {
	// Пример запуска обновлений
	go func() {
		for {
			h.UpdTicker()
			log.Print("upd")
			time.Sleep(1 * time.Minute)
		}
	}()
}
