package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/swastik-gautam/url-shortener/cache"
	"github.com/swastik-gautam/url-shortener/store"
)

type Handler struct {
	store *store.Store
}

func New(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if body.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	shortCode, err := h.store.Set(body.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"short_code": shortCode,
	})
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // remove leading "/"

	longURL, found := cache.Get(shortCode)
	if found {
		http.Redirect(w, r, longURL, http.StatusFound)
		return
	}

	longURL, err := h.store.Get(shortCode)
	if err != nil {
		http.Error(w, "short  URL not found", http.StatusNotFound)
		return
	}

	cache.Set(shortCode, longURL)

	http.Redirect(w, r, longURL, http.StatusFound)
}
