package usecase

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"main.go/cmd/internal/repository"
)

// поиск по url
func GetFetch(c echo.Context) error {
	NameTicker := c.Param("ticker")

	// проверка данных в бд

	exists, err := CheckTicker(NameTicker) // ?
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Database error")
	}
	if !exists {
		return ErrorResponse(c, http.StatusNotFound, "Ticker not found")
	}

	// time весь переход надо сделать в отдельном файле!!
	startUnix, endUnix, err := transition(c, "date_from", "date_to")
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err)
	}
	// fmt.Println(startUnix)
	// fmt.Println(endUnix)

	//Смотрим в бд
	////////////////////////// сделать все в репозитории
	flagDataFrom := true
	var startPrice tickerHistory // не должно быть
	if err := db.DB.Where("ticker=? AND timestamp = ?", NameTicker, startUnix).First(&startPrice).Error; err != nil {
		log.Println("нет даты в бд(data_from)")
		flagDataFrom = false
	}
	flagDateTo := true
	var endPrice tickerHistory // не должно быть
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
	/////////////////////////////////
	//выбить в отдельную функцию

	percDiffStr := calculatorDiff(endPrice.Price, startPrice.Price)
	//
	return c.JSON(http.StatusOK, priceDiff{
		Name:       NameTicker,
		Price:      endPrice.Price,
		Difference: percDiffStr,
	})

}

// добавления тикера
func PostAddTicker(c echo.Context) error {
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
	service := repository.BinanceClientPost()
	Resp, err := service.Symbol(ticker.Ticker).Do(context.Background())
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Error, 400")
	}
	price, err := strconv.ParseFloat(Resp[0].Price, 64)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Error 500")
	}

	//добавления данных

	//??
	NewTicker := tickerDB{
		Ticker: ticker.Ticker,
		Price:  price,
	}
	//SaveTicker
	if err := db.DB.Create(&NewTicker).Error; err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Error 500.\nError when adding")
	}

	//??
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
