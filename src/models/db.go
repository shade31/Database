package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Database *pgxpool.Pool

func OpenDatabaseConnection() {
	var err error
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Douala", host, username, password, databaseName, port)

	// Create database connection
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Ошибка при получении соединения из пула БД")
	}

	Database = pool

	fmt.Println("Соединение к БД установлено!")
}

func CloseDatabaseConnection() {
	if Database != nil {
		Database.Close()
		fmt.Println("Соединение с БД закрыто.")
	}
}
