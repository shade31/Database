package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config())
	if err != nil {
		log.Fatal("Ошибка при установке соединения с БД")
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Ошибка при получении соединения из пула БД")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Не удалось подключиться к БД")
	}

	fmt.Println("Соединение к БД установлено!")

	// Database queries
	InsertQuery(connPool, 5, "Chai", 100, 100)

	defer connPool.Close()

}

func InsertQuery(p *pgxpool.Pool, id int, name string, price int, count int) {
	_, err := p.Exec(context.Background(), "insert into products(id, name, price, count) values($1, $2, $3, $4)", id, name, price, count)
	if err != nil {
		log.Fatal("Ошибка при записи данных в таблицу")
	}
}

func UpdateQuery(p *pgxpool.Pool, id int, price int, count int) {
	_, err := p.Exec(context.Background(), "update products set price = $2, count = $3 where id = $1", id, price, count)
	if err != nil {
		log.Fatal("Ошибка при обновлении данных в таблице")
	}
}

func SelectQuery(p *pgxpool.Pool) {
	rows, err := p.Query(context.Background(), "select * from products")
	if err != nil {
		log.Fatal("Ошибка при получении данных с табилцы")
	}

	for rows.Next() {
		var id int
		var name string
		var price int
		var count int
		err = rows.Scan(&id, &name, &price, &count)
		if err != nil {
			log.Fatal("Ошибка при получении данных с табилцы")
		}
		fmt.Printf("id: %d, name: %s, price: %d, count: %d\n", id, name, price, count)
	}
}

func DeleteQuery(p *pgxpool.Pool, id int) {
	_, err := p.Exec(context.Background(), "delete from products where id = $1", id)
	if err != nil {
		log.Fatal("Ошибка при удалении записи из таблицы")
	}
}

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	// Your own Database URL
	const DATABASE_URL string = "postgres://postgres:123@localhost:5432/Shop"

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}
