package deliv

import (
	"fmt"
	"log"
	"net/http"
	"tass-binance/internal/module"
	"tass-binance/internal/module/repository/pstgrs"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	us module.UseCase
}
type Message struct {
	Message string `json:"message"`
}

type responseDiff struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Difference string  `json:"difference"`
}

func NewHandler(us module.UseCase) *Handler {
	return &Handler{
		us: us,
	}
}

func (h *Handler) AddTicker(c echo.Context) error {
	ticker, err := pstgrs.BindTicker(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Message{Message: fmt.Sprintf("error: Invalid request %s", err)})
	}
	err = h.us.ProcessTicker(ticker)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Message{Message: fmt.Sprintf("error, %s", err)})
	}
	return c.JSON(http.StatusOK, Message{Message: "ticker added successfully"})
}

// GetFetch
func (h *Handler) TickerDiff(c echo.Context) error {
	log.Print("start TickerDiff")
	// взять данные из url
	ticker, dateFrom, dateTo, err := pstgrs.ParamTicker(c)
	if err != nil {
		return err
	}
	log.Print("ParamTicker successfully")
	// отдать данные в процесс - получить разницу и ошибку
	dif, priceTo, err := h.us.ProcessTickerDiff(ticker, dateFrom, dateTo)
	// проверить ошибку
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, Message{Message: "error..."})
	}
	//вывести данные, построим структуру
	return c.JSON(http.StatusOK, responseDiff{
		Name:       ticker,
		Price:      priceTo,
		Difference: fmt.Sprintf("%.2f%%", dif),
	})
}

func (h *Handler) UpdTicker() error {
	err := h.us.TickerUpdateProcess()
	if err != nil {
		return fmt.Errorf("error: Invalid request %s", err)
	}
	return nil
}
