package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swastik-gautam/url-shortener/encoder"
)

type Store struct {
	db *pgxpool.Pool
}

func New(connStr string) (*Store, error) {
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Set(longURL string) (string, error) {
	var id int
	err := s.db.QueryRow(
		context.Background(),
		"INSERT INTO urls (long_url) VALUES ($1) RETURNING id",
		longURL,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	shortCode := encoder.Encode(id)

	_, err = s.db.Exec(
		context.Background(),
		"UPDATE urls SET short_code = $1 WHERE id = $2",
		shortCode, id,
	)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *Store) Get(shortCode string) (string, error) {
	var longURL string
	err := s.db.QueryRow(
		context.Background(),
		"SELECT long_url FROM urls WHERE short_code = $1",
		shortCode,
	).Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
