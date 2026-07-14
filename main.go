package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swastik-gautam/url-shortener/handlers"
	"github.com/swastik-gautam/url-shortener/store"
)

func main() {
	connStr := "postgres://swastik:swastik@27@localhost:5432/urlshortener"

	s, err := store.New(connStr)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}

	h := handlers.New(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", h.Shorten)
	mux.HandleFunc("/", h.Redirect)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
