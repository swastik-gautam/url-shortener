package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/swastik-gautam/url-shortener/cache"
	"github.com/swastik-gautam/url-shortener/handlers"
	"github.com/swastik-gautam/url-shortener/store"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env vars")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	s, err := store.New(connStr)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}

	// just connect to Redis (uses the global RDB)
	cache.ConnectRedis()

	h := handlers.New(s) // only pass the store

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", h.Shorten)
	mux.HandleFunc("/", h.Redirect)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
