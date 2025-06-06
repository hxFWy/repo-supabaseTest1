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
	err := r.db.QueryRowContext(context.Background(), "SELECT * FROM public.users WHERE username = $1", username).
		Scan(&fetchedUser.Id,
			&fetchedUser.Created_at,
			&fetchedUser.Username,
			&fetchedUser.Password)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch user with username: %s - %v", username, err)
	}

	return &fetchedUser, nil
}

func (r *Repository) GetUserById(id int) (*types.User, error) {

	var fetchedUser types.User
	err := r.db.QueryRowContext(context.Background(), "SELECT * FROM public.users WHERE id = $1", id).
		Scan(&fetchedUser.Id,
			&fetchedUser.Created_at,
			&fetchedUser.Username,
			&fetchedUser.Password)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch user with id: %d - %v", id, err)
	}

	return &fetchedUser, nil
}

func (r *Repository) CreateUserTx(ctx context.Context, tx *sql.Tx, payload types.RegisterUserPayload) (*types.User, error) {

	var newUser types.User

	err := tx.QueryRowContext(ctx, `INSERT INTO public.users (username, password)
														VALUES ($1, $2)
														RETURNING id, created_at, username`,

		payload.Username, payload.Password).Scan(
		&newUser.Id,
		&newUser.Created_at,
		&newUser.Username,
	)

	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
