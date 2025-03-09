package app

import (
	"tass-binance/infrastructure/database"
	"tass-binance/internal/module/deliv"
	"tass-binance/internal/module/repository/extrenal_api"
	"tass-binance/internal/module/repository/pstgrs"
	"tass-binance/internal/module/usecase"
)

func InitApp(server *Server) error {
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
	tickerUseCase := usecase.NewUseCase(repo, extranalAPI)

	// Инициализация обработчиков HTTP-запросов (deliv) и передача в них usecase
	handler := deliv.NewHandler(tickerUseCase)

	// Регистрация маршрутов в сервере
	deliv.MapRoutes(server.echo, *handler)

	return nil
}
