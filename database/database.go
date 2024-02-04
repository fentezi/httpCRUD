package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB_URL := os.Getenv("DB_URL")
	DB, err = sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal(err)
	}

}
