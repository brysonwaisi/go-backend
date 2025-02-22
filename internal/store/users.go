package store

import (
	"context"
	"database/sql"

	// "github.com/go-chi/chi/v5"
	// "github.com/brysonwaisi/go-backend/internal/store"
	// "github.com/lib/pq"
)

// this is the model
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password int64 `json:"_"`
	CreatedAt string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (username, email, password, tags)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx, 
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan (
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}