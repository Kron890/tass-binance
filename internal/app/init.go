package app

import (
	"tass-binance/infrastructure/database"
	"tass-binance/internal/module/repository/extrenal_api"
	"tass-binance/internal/module/repository/pstgrs"
)

func (s *Server) initApp() error {
	//подключения к бд
	dbConnect, err := database.NewDbConnection()
	if err != nil {
		return err
	}

	// Создание репозитория для работы с тикерами через PostgreSQL
	repo := pstgrs.NewRepo(dbConnect)

	// Инициализация клиента для взаимодействия с внешним API
	extranalAPI := extrenal_api.NewApi()

	// Создание слоя бизнес-логики (usecase) и передача в него зависимостей
	tickerUseCase := 
}
