package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swastik-gautam/url-shortener/handlers"
	"github.com/swastik-gautam/url-shortener/store"
)

func main() {
	s := store.New()
	h := handlers.New(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", h.Shorten)
	mux.HandleFunc("/", h.Redirect)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
