package deliv

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func MapRoutes(e *echo.Echo, h Handler) {
	e.POST("/add_ticker", h.AddTicker)
	//добавить просморт тикеров в бд
	// e.GET("/fetch/:ticker",)

}

func StartUpd(c *echo.Echo, h Handler) {
	for {
		h.UpdTicker()
		fmt.Println("Upd")
		time.Sleep(60 * time.Second)
	}
}
