package store

import (
	"sync"

	"github.com/swastik-gautam/url-shortener/encoder"
)

type Store struct {
	mu      sync.Mutex
	urls    map[string]string
	counter int
}

func New() *Store {
	return &Store{
		urls: make(map[string]string),
	}
}

func (s *Store) Set(longURL string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++
	shortCode := encoder.Encode(s.counter)
	s.urls[shortCode] = longURL
	return shortCode
}

func (s *Store) Get(shortCode string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	longURL, ok := s.urls[shortCode]
	return longURL, ok
}
