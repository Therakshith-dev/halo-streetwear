package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dsn := "host=localhost port=5432 user=rakshiths password=" +
		os.Getenv("DB_PASSWORD") +
		" dbname=myapp_db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
	return db
}
