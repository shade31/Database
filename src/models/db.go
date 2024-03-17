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

// func InsertQuery(p *pgxpool.Pool, id int, name string, price int, count int) {
// 	_, err := p.Exec(context.Background(), "insert into products(id, name, price, count) values($1, $2, $3, $4)", id, name, price, count)
// 	if err != nil {
// 		log.Fatal("Ошибка при записи данных в таблицу")
// 	}
// }

func UpdateQuery(p *pgxpool.Pool, id int, price int, count int) {
	_, err := p.Exec(context.Background(), "update products set price = $2, count = $3 where id = $1", id, price, count)
	if err != nil {
		log.Fatal("Ошибка при обновлении данных в таблице")
	}
}

// func SelectQuery(p *pgxpool.Pool) {
// 	rows, err := p.Query(context.Background(), "select * from products")
// 	if err != nil {
// 		fmt.Println(err)
// 		log.Fatal("Ошибка при получении данных с таблицы")
// 	}

// 	for rows.Next() {
// 		var id int
// 		var name string
// 		var price int
// 		var count int
// 		err = rows.Scan(&id, &name, &price, &count)
// 		if err != nil {
// 			fmt.Println(err)
// 			log.Fatal("Ошибка при получении данных с таблицы - 2")
// 		}
// 		fmt.Printf("id: %d, name: %s, price: %d, count: %d\n", id, name, price, count)
// 	}
// }

func DeleteQuery(p *pgxpool.Pool, id int) {
	_, err := p.Exec(context.Background(), "delete from products where id = $1", id)
	if err != nil {
		log.Fatal("Ошибка при удалении записи из таблицы")
	}
}

// func Config() *pgxpool.Config {
// 	const defaultMaxConns = int32(4)
// 	const defaultMinConns = int32(0)
// 	const defaultMaxConnLifetime = time.Hour
// 	const defaultMaxConnIdleTime = time.Minute * 30
// 	const defaultHealthCheckPeriod = time.Minute
// 	const defaultConnectTimeout = time.Second * 5

// 	// Your own Database URL
// 	const DATABASE_URL string = "postgres://postgres:123@localhost:5432/Shop"

// 	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
// 	if err != nil {
// 		log.Fatal("Failed to create a config, error: ", err)
// 	}

// 	dbConfig.MaxConns = defaultMaxConns
// 	dbConfig.MinConns = defaultMinConns
// 	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
// 	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
// 	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
// 	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

// 	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
// 		log.Println("Before acquiring the connection pool to the database!!")
// 		return true
// 	}

// 	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
// 		log.Println("After releasing the connection pool to the database!!")
// 		return true
// 	}

// 	dbConfig.BeforeClose = func(c *pgx.Conn) {
// 		log.Println("Closed the connection pool to the database!!")
// 	}

// 	return dbConfig
// }
