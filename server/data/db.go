package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(ctx context.Context, dbURL string) (*Store, error) {
	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	} 
	
	return &Store{
		db: dbpool,
	}, nil
}

func (s *Store) Ping(ctx context.Context) error {
	// check database connection
	if err := s.db.Ping(ctx); err != nil {
		return fmt.Errorf("database connection error: %w", err)
	}

	return nil
}

func (s *Store) Close() {
	if s.db != nil {
		s.db.Close()
	}
}