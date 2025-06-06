package player

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

func (r *Repository) GetPlayerByUserId(id int) (*types.Player, error) {

	var fetchedPlayer types.Player

	err := r.db.QueryRowContext(context.Background(), "SELECT * FROM public.players WHERE user_id = $1", id).
		Scan(&fetchedPlayer.User_id,
			&fetchedPlayer.Money,
			&fetchedPlayer.Position,
			&fetchedPlayer.Stamina,
			&fetchedPlayer.Skill,
			&fetchedPlayer.Created_at)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch player by user id: %d - %v", id, err)
	}

	return &fetchedPlayer, nil
}

func (r *Repository) CreatePlayerTx(ctx context.Context, tx *sql.Tx, payload types.CreatePlayerPayload) error {
	err := tx.QueryRowContext(ctx, "INSERT INTO public.players (user_id, position) VALUES ($1, $2)",
		payload.User_id, payload.Position)

	if err != nil {
		return err.Err()
	}
	return nil
}
