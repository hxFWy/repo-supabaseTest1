package user

import (
	"context"
	"database/sql"
	"fmt"

	"supabase-testProject1/internal/types"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) GetUserByUsername(username string) (*types.User, error) {

	var fetchedUser types.User
	err := r.db.QueryRowContext(context.Background(), "SELECT * FROM users WHERE username = $1", username).Scan(&fetchedUser)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch user with username: %s - %v", username, err)
	}

	return &fetchedUser, nil
}

func (r *Repository) GetUserById(id int) (*types.User, error) {

	var fetchedUser types.User
	err := r.db.QueryRowContext(context.Background(), "SELECT * FROM users WHERE id = $1", id).Scan(&fetchedUser)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch user with id: %d - %v", id, err)
	}

	return &fetchedUser, nil
}

func (r *Repository) CreateUser(payload types.RegisterUserPayload) error {
	err := r.db.QueryRowContext(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)",
		payload.Username, payload.Password)
	if err != nil {
		return err.Err()
	}
	return nil
}
