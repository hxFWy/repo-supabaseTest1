package training

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) TrainPlayerById(user_id int) error {

	_, err := r.db.ExecContext(context.Background(),
		`UPDATE public.players
		SET skill = skill + 1
		WHERE user_id = $1`, user_id)

	if err != nil {
		return fmt.Errorf("could not update skill: %w", err)
	}

	return nil
}
