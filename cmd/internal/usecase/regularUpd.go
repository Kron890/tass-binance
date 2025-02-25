package usecase

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (db *DataBase) RegularUpd() {
	tickerService := BinacnClientRegular()
	for {
		var tickers []tickerDB

		if err := db.DB.Select("ticker").Find(&tickers).Error; err != nil {
			log.Println("Error receiving tickers:", err)
		}

		for _, name := range tickers {
			priceResp, err := tickerService.Symbol(name.Ticker).Do(context.Background())
			if err != nil {
				log.Println("Ошибка, данные не отправились")

			}
			price, err := strconv.ParseFloat(priceResp[0].Price, 64)
			if err != nil {
				log.Fatal(err)
			}
			if err := db.DB.Model(&tickerDB{}).Where("ticker = ?", name.Ticker).Update("price", price).Error; err != nil {
				log.Println("Ошибка, в добавлении данных")
			}

			tickerHistory := tickerHistory{
				Ticker:    name.Ticker,
				Price:     price,
				Timestamp: time.Now().Unix(),
			}
			if err := db.DB.Create(&tickerHistory).Error; err != nil {
				log.Println("Ошибка, в созданине данных")
			}
			fmt.Println("Данные обновлены")
			time.Sleep(1 * time.Minute)
		}

	}

}
