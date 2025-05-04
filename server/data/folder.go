package data

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Folder struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Store) CreateFolder(ctx context.Context, userID string, name string) (*Folder, error) {
	var f Folder
	
	query := `INSERT INTO folders (user_id, name) 
			  VALUES ($1, $2) 
			  RETURNING id, user_id, name, created_at`
	
	err := s.db.QueryRow(ctx, query, userID, name).Scan(&f.ID, &f.UserID, &f.Name, &f.CreatedAt)
	if err != nil {
		log.Printf("Query execution error: %v", err)
		return nil, fmt.Errorf("failed to create folder: %w", err)
	}

	return &f, nil
}

func (s *Store) DeleteFolder(ctx context.Context, userID string, name string) error {
	query := `DELETE FROM folders WHERE user_id = $1 AND name = $2`

	result, err := s.db.Exec(ctx, query, userID, name)
	if err != nil {
		return fmt.Errorf("failed to delete folder: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no folder found with name: %s for user: %s", name, userID)
	}

	return nil
}