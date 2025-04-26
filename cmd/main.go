package main

import (
	"auth-service/database"
	"auth-service/internal/server"
)

func main() {
	err := database.Connect()

	if err != nil {
		panic(err)
	}

	err = server.Start()

	if err != nil {
		panic(err)
	}
}
