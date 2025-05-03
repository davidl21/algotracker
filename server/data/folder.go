package data

import (
	"context"
	"fmt"
	"log"
)

type Folder struct {
	ID string
	UserID string
	Name string
	CreatedAt string 
}

func (s *Store) CreateFolder(ctx context.Context, userID string, name string) (*Folder, error) {
	var f Folder
	
	// Add connection check
	if err := s.db.Ping(ctx); err != nil {
		log.Printf("Database connection check failed: %v", err)
		return nil, fmt.Errorf("database connection error: %w", err)
	}
	
	query := `INSERT INTO folders (user_id, name) 
			  VALUES ($1, $2) 
			  RETURNING id, user_id, name, created_at`
	
	log.Printf("Executing query with userID: %s, name: %s", userID, name)
	
	err := s.db.QueryRow(ctx, query, userID, name).Scan(&f.ID, &f.UserID, &f.Name, &f.CreatedAt)
	if err != nil {
		log.Printf("Query execution error: %v", err)
		return nil, fmt.Errorf("failed to create folder: %w", err)
	}

	return &f, nil
}