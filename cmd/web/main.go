package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Kaveh-Goodarzi/url-shortner/internal/config"
	"github.com/Kaveh-Goodarzi/url-shortner/internal/handlers"
	"github.com/jackc/pgx/v5"
)

func main() {
	dbURL, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("unable to connect to database")
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("database connection failed")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /urls", handlers.CreateURLHandler)
	mux.HandleFunc("GET /{code}", handlers.RedirectURLHandler)

	log.Println("starting server on localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
