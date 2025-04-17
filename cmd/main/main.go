package main

import (
	"errors"
	"net/http"
	"tass-binance/config"
	"tass-binance/internal/app"
	"tass-binance/pkg/logger"
)

func main() {
	// представляет собой веб-сервер
	l := logger.NewLogger()
	server := app.NewServer()

	configuration, err := config.GetConfig()
	if err != nil {
		l.Errorf("Failed to get config: %s", err.Error())
		panic(err)
	}
	err = app.InitApp(server, configuration, l)
	if err != nil {
		l.Errorf("Initialization error: %s", err)
		panic(err)
	}
	l.Infof("Server successfully init")

	err = server.Start()
	if errors.Is(err, http.ErrServerClosed) {
		l.Errorf("server two closed\n")
	} else if err != nil {
		l.Errorf("server error: %v\n", err)
		panic(err)
	}
}
