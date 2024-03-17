package models

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count int    `json:"count"`
}

func SelectQuery(p *pgxpool.Pool) (string, error) {
	var result string
	rows, err := p.Query(context.Background(), "select * from products")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Ошибка при получении данных с таблицы")
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price int
		var count int
		err = rows.Scan(&id, &name, &price, &count)
		if err != nil {
			fmt.Println(err)
			log.Fatal("Ошибка при получении данных с таблицы - 2")
			return "", err
		}
		result += fmt.Sprintf("id: %d, name: %s, price: %d, count: %d |", id, name, price, count)
	}

	return result, nil
}

func SelectQueryById(p *pgxpool.Pool, id int) (string, error) {
	var result string
	rows, err := p.Query(context.Background(), "select * from products where id=$1", id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Ошибка при получении данных с таблицы")
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price int
		var count int
		err = rows.Scan(&id, &name, &price, &count)
		if err != nil {
			fmt.Println(err)
			log.Fatal("Ошибка при получении данных с таблицы - 2")
			return "", err
		}
		result += fmt.Sprintf("id: %d, name: %s, price: %d, count: %d |", id, name, price, count)
	}

	return result, nil
}

func InsertQuery(p *pgxpool.Pool, product Product) error {
	_, err := p.Exec(context.Background(), "insert into products(id, name, price, count) values($1, $2, $3, $4)", product.Id, product.Name, product.Price, product.Count)
	if err != nil {
		log.Fatal("Ошибка при записи данных в таблицу", err)
		return err
	}
	fmt.Println("Новая запись успешно добавлена в таблицу products")
	return nil
}
