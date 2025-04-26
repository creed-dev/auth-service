package main

import (
	"auth-service/database"
	"auth-service/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := database.Connect()

	if err != nil {
		log.Println("Приложение остановлено")
		panic(fmt.Sprintf("Ошибка при подключении к базе данных:\n%s", err))
	}

	defer database.Close()

	go func() {
		err = server.Start()

		if err != nil {
			log.Println("Приложение остановлено")
			database.Close()
			panic(fmt.Sprintf("Ошибка при старте сервера:\n%s", err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Приложение остановлено")
}
