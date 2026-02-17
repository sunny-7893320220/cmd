package main

import (
	"auth/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	name      string
	price     float64
	avaliable bool
}

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg.Port)
	fmt.Println(cfg.RedisPort)

	connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)

	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	createTable(db)
	log.Println("Table created")

	product := []Product{{"Laptop", 1000.00, true}, {"Mobile", 500.00, true}, {"Tablet", 200.00, true}, {"Laptop", 1000.00, true}, {"Mobile", 500.00, true}, {"Tablet", 200.00, true}}

	fmt.Println(len(product))
	for _, product := range product {
		pk := insertProduct(db, product)
		fmt.Println(pk)
	}
	log.Println("Product inserted")
}

func createTable(db *sql.DB) {
	query :=
		`CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		avaliable BOOLEAN NOT NULL,
		created timestamp DEFAULT now()
		)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product 
	(name, price, avaliable)
	VALUES ($1, $2, $3) RETURNING id`
	var pk int
	err := db.QueryRow(query, product.name, product.price, product.avaliable).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
