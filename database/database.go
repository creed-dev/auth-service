package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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

	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		return fmt.Errorf("ошибка открытия подключения к базе данных\n%s", err)
	}

	err = db.Ping()

	if err != nil {
		return fmt.Errorf("ошибка пинга базы даных\n%s", err)
	}

	fmt.Println("Успешное подключение к базе данных")

	return nil
}
