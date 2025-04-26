package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	err := godotenv.Load()

	if err != nil {
		return fmt.Errorf("ошибка загрузки .env файла")
	}

	connectionStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err = sql.Open("postgres", connectionStr)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	err = db.Ping()

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	log.Println("Успешное соединение с базой данных")

	return nil
}

func Close() {
	if db == nil {
		return
	}

	err := db.Close()

	if err != nil {
		log.Printf("Ошибка при закрытии соединения с базой данных:\n%s", err)
	} else {
		log.Println("Соединение с базой данных закрыто")
	}
}

func GetDatabase() *sql.DB {
	return db
}
