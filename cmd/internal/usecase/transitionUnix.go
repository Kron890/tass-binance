package usecase

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

// как будут работать ошибки ?
func transition(c echo.Context, date_from, date_to string) (int64, int64, error) {
	dateFromParam := c.Param(date_from)
	loc, _ := time.LoadLocation("Europe/Moscow") // Загружаем московский часовой пояс
	dateFromTime, err := time.ParseInLocation("02.01.2006 15:04:05", dateFromParam, loc)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid data provided(date_from)")
	}
	dateToParam := c.Param(date_to)
	dateToTime, err := time.ParseInLocation("02.01.2006 15:04:05", dateToParam, loc)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid data provided(date_to)")
	}

	return dateFromTime.Unix(), dateToTime.Unix(), nil
}
