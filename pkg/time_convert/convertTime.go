package time_convert

import (
	"fmt"
	"time"
)

func ConvertTime(dateFrom, dateTo string) (int64, int64, error) {
	loc, _ := time.LoadLocation("Europe/Moscow") // Загружаем московский часовой пояс
	dateFromTime, err := time.ParseInLocation("02.01.2006 15:04:05", dateFrom, loc)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid data provided(date_from)")
	}
	dateToTime, err := time.ParseInLocation("02.01.2006 15:04:05", dateTo, loc)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid data provided(date_to)")
	}

	return dateFromTime.Unix(), dateToTime.Unix(), nil
}
