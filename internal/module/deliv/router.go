package deliv

import (
	"github.com/labstack/echo/v4"
)

func MapRoutes(e *echo.Echo, h Handler) {
	e.POST("/add_ticker", h.AddTicker)
	//добавить просморт тикеров в бд
	// e.GET("/fetch/:ticker",)

}
