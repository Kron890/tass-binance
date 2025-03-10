package deliv

import (
	"fmt"
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

func (h *Handler) GetTickerDiff(c echo.Context) error {
	// взять данные из url
	// отдать данные в процесс - получить разницу и ошибку
	// проверить ошибку
	//вывести данные, построим структуру
	return nil
}

func (h *Handler) UpdTicker() error {
	err := h.us.TickerUpdateProcess()
	if err != nil {
		return fmt.Errorf("error: Invalid request %s", err)
	}
	return nil
}
