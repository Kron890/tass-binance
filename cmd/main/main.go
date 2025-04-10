package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"tass-binance/config"
	"tass-binance/internal/app"
)

func main() {
	// представляет собой веб-сервер
	server := app.NewServer()

	configuration, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = app.InitApp(server, configuration)
	if err != nil {
		log.Print("Initialization error: ", err)
		panic(err)
	}
	log.Print("Server successfully init")

	err = server.Start()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server two closed\n")
	} else if err != nil {
		fmt.Printf("server error: %v\n", err)
		os.Exit(1)
	}
}
