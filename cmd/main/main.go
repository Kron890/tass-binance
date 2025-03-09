package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"tass-binance/internal/app"
)

func main() {
	//представляет собой веб-сервер
	server := app.NewServer()

	err := app.InitApp(server)
	if err != nil {
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
