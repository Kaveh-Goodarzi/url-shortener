package main

import (
	"log"
	"net/http"

	"github.com/Kaveh-Goodarzi/url-shortner/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /urls", handlers.CreateURLHandler)

	log.Println("starting server on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
