package usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// поиск по url
func (db *DataBase) GetFetch(c echo.Context) error {

	NameTicker := c.Param("ticker")
	// проверка данных в бд
	var count int64
	if err := db.DB.Model(tickerDB{}).Where("ticker=?", NameTicker).Count(&count).Error; err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Database error")
	}
	if count == 0 {
		return ErrorResponse(c, http.StatusNotFound, "Tiker not found")
	}
	// time весь переход надо сделать в отдельном файле!!
	dateFromParam := c.Param("date_from")
	loc, _ := time.LoadLocation("Europe/Moscow") // Загружаем московский часовой пояс
	dateFromTime, err := time.ParseInLocation("02.01.2006 15:04:05", dateFromParam, loc)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid data provided(date_from)")
	}
	dateToParam := c.Param("date_to")
	dateToTime, err := time.ParseInLocation("02.01.2006 15:04:05", dateToParam, loc)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid data provided(date_to)")
	}
	//UNIX
	startUnix := dateFromTime.Unix()
	endUnix := dateToTime.Unix()

	// fmt.Println(startUnix)
	// fmt.Println(endUnix)

	//Смотрим в бд
	flagDataFrom := true
	var startPrice tickerHistory
	if err := db.DB.Where("ticker=? AND timestamp = ?", NameTicker, startUnix).First(&startPrice).Error; err != nil {
		log.Println("нет даты в бд(data_from)")
		flagDataFrom = false
	}
	flagDateTo := true
	var endPrice tickerHistory
	if err := db.DB.Where("ticker = ? AND timestamp = ?", NameTicker, endUnix).First(&endPrice).Error; err != nil {
		log.Println("нет даты в бд(data_to)")
		flagDateTo = false
	}
	// type DifferenceTicker PriceDiff
	if !flagDataFrom || !flagDateTo {
		return c.JSON(http.StatusOK, priceDiff{
			Name:       NameTicker,
			Price:      0,
			Difference: "0%",
		})
	}

	//выбить в отдельную функцию

	percDiff := ((endPrice.Price - startPrice.Price) / startPrice.Price) * 100
	percDiffStr := fmt.Sprintf("%.2f%%", percDiff)
	//
	return c.JSON(http.StatusOK, priceDiff{
		Name:       NameTicker,
		Price:      endPrice.Price,
		Difference: percDiffStr,
	})

}

// добавления тикера
func (db *DataBase) PostAddTicker(c echo.Context) error {
	var ticker tickerDB
	//вытаскиваем из запорса
	if err := c.Bind(&ticker); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Error, incorrect data form")
	}

	//проверяем есть ли в бд данные
	var count int64
	if err := db.DB.Model(tickerDB{}).Where("ticker=?", ticker.Ticker).Count(&count).Error; err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Database error")
	}
	if count != 0 {
		return ErrorResponse(c, http.StatusConflict, "The data already exists")
	}

	// Вытаскиваем из бинанса данные
	service := BinanceClientPost()
	Resp, err := service.Symbol(ticker.Ticker).Do(context.Background())
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Error, 400")
	}
	price, err := strconv.ParseFloat(Resp[0].Price, 64)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Error 500")
	}

	//добавления данных

	//оставить
	NewTicker := tickerDB{
		Ticker: ticker.Ticker,
		Price:  price,
	}
	//SaveTicker
	if err := db.DB.Create(&NewTicker).Error; err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Error 500.\nError when adding")
	}

	//оставить
	NewTickerHistory := tickerHistory{
		Ticker:    ticker.Ticker,
		Price:     price,
		Timestamp: time.Now().Unix(),
	}
	// SaveTicker
	if err := db.DB.Create(&NewTickerHistory).Error; err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Error 500.\nError when adding")
	}

	return c.JSON(http.StatusOK, Status{Message: "Ticker added"})
}
