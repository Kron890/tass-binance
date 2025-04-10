package deliv

import (
	"fmt"
	"log"
	"net/http"
	"tass-binance/internal/module"
	"tass-binance/internal/module/entity"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	us module.UseCase
}

func NewHandler(us module.UseCase) *Handler {
	return &Handler{
		us: us,
	}
}

func (h *Handler) AddTicker(c echo.Context) error {
	var ticker entity.TickerAddRequest
	err := c.Bind(&ticker)
	if err != nil {
		return err
	}
	err = h.us.AddTicker(ticker)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Message{Message: fmt.Sprintf("error, %s", err)})
	}
	return c.JSON(http.StatusOK, entity.Message{Message: "ticker added successfully"})
}

// GetFetch
func (h *Handler) TickerDiff(c echo.Context) error {
	var ticker entity.TickerDiffRequest
	ticker.Name = c.Param("ticker")
	if ticker.Name == "" {
		return fmt.Errorf("incorrect request data (ticker)")
	}
	ticker.DateFrom = c.Param("date_from")
	if ticker.DateFrom == "" {
		fmt.Println("dateFrom")
	}
	ticker.DateTo = c.Param("date_to")
	if ticker.DateTo == "" {
		fmt.Println("dateFrom")
	}
	result, err := h.us.ProcessTickerDiff(ticker)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, entity.Message{Message: "error..."})
	}
	return c.JSON(http.StatusOK, result)
}
