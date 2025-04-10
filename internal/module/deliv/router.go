package deliv

import (
	"github.com/labstack/echo/v4"
)

func MapRoutes(e *echo.Echo, h Handler) {

	e.POST("/add_ticker", h.AddTicker)
	e.GET("/fetch/:ticker/:date_from/:date_to", h.TickerDiff)
}
