package repositories

import (
	"database/sql"
	"fmt"
	"mongo_connection/models"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// DSN format: "user:password@tcp(127.0.0.1:3306)/dbname"
	dsn := "app_developer:1234@tcp(127.0.0.1:3306)/production_db"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to MySQL")
}

func GetUsers() ([]models.User, error) {
	rows, err := DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
