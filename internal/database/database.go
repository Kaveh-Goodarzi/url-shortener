package database

import (
	"database/sql"
	"log"

	"github.com/Kaveh-Goodarzi/url-shortner/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var URLstore = make(map[string]string)

var DB *sql.DB

func InitDB() {
	dbURL, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("unable to connect to database")
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("database connection failed")
	}
	// defer DB.Close()
}
