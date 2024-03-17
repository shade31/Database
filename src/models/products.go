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

func SelectQuery(p *pgxpool.Pool) ([]Product, error) {
	var products []Product
	rows, err := p.Query(context.Background(), "select * from products")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Ошибка при получении данных с таблицы")
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Count)
		if err != nil {
			fmt.Println(err)
			log.Fatal("Ошибка при получении данных с таблицы")
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func SelectQueryById(p *pgxpool.Pool, id int) ([]Product, error) {
	var products []Product
	rows, err := p.Query(context.Background(), "select * from products where id=$1", id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Ошибка при получении данных с таблицы")
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Count)
		if err != nil {
			fmt.Println(err)
			log.Fatal("Ошибка при получении данных с таблицы")
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func InsertQuery(p *pgxpool.Pool, product Product) error {
	_, err := p.Exec(context.Background(), "insert into products(name, price, count) values($1, $2, $3)", product.Name, product.Price, product.Count)
	if err != nil {
		log.Fatal("Ошибка при записи данных в таблицу", err)
		return err
	}
	fmt.Println("Новая запись успешно добавлена в таблицу products")
	return nil
}

func UpdateQuery(p *pgxpool.Pool, id int, price int, count int) error {
	_, err := p.Exec(context.Background(), "update products set price = $2, count = $3 where id = $1", id, price, count)
	if err != nil {
		log.Fatal("Ошибка при обновлении данных в таблице", err)
		return err
	}
	return nil
}

func DeleteQuery(p *pgxpool.Pool, id int) error {
	_, err := p.Exec(context.Background(), "delete from products where id = $1", id)
	if err != nil {
		log.Fatal("Ошибка при удалении записи из таблицы")
	}

	return nil
}
