package app

import (
	"tass-binance/config"
	"tass-binance/infrastructure/database"
	"tass-binance/internal/module/deliv"
	"tass-binance/internal/module/repository/extrenal_api"
	"tass-binance/internal/module/repository/pstgrs"
	"tass-binance/internal/module/usecase"
	"tass-binance/pkg/logger"
)

func InitApp(server *Server, config config.Config, l *logger.Logger) error {
	//подключения к бд

	dbConnect, err := database.NewDbConnection(config)
	if err != nil {
		return err
	}
	l.Infof("successful connection to the database")

	// Создание репозитория для работы с тикерами через PostgreSQL
	repo := pstgrs.NewRepo(dbConnect)
	// Инициализация клиента для взаимодействия с внешним API
	extrenalAPI := extrenal_api.NewApi()
	extrenalAPI.Init()
	l.Infof("successful connection to the database")

	// Создание слоя бизнес-логики (usecase) и передача в него зависимостей
	tickerUseCase := usecase.NewUseCase(repo, extrenalAPI, l)

	// Инициализация обработчиков HTTP-запросов (deliv) и передача в них usecase
	handler := deliv.NewHandler(tickerUseCase)

	// Регистрация маршрутов в сервере
	deliv.MapRoutes(server.echo, *handler)

	return nil
}
