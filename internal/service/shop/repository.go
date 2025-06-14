package shop

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

func (r *Repository) GetItemsList(ctx context.Context) ([]*types.Item, error) {
	items := make([]*types.Item, 0)

	const q = `
        SELECT
            id,
            name,
            slot,
            cost,
            skill_bonus
        FROM public.items
    `

	rows, err := r.db.QueryContext(ctx, q)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve the items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := new(types.Item)
		if err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Slot,
			&item.Cost,
			&item.Skill_bonus,
		); err != nil {
			return nil, fmt.Errorf("could not create item object with retrieved data: %w", err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return items, nil
}

func (r *Repository) GetItemById(ctx context.Context, itemId int) (*types.Item, error) {
	item := new(types.Item)

	const q = `
		SELECT
			id,
			name,
			slot,
			cost,
			skill_bonus
		FROM public.items
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, q, itemId).Scan(
		&item.Id,
		&item.Name,
		&item.Slot,
		&item.Cost,
		&item.Skill_bonus,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item with id %d not found", itemId)
		}
		return nil, fmt.Errorf("could not retrieve item: %w", err)
	}

	return item, nil
}

func (r *Repository) AddItemToPlayerTx(ctx context.Context, tx *sql.Tx, userId, itemId int) error {
	const q = `
		INSERT INTO public.player_items (player_id, item_id)
		VALUES ($1, $2)
	`

	_, err := tx.ExecContext(ctx, q, userId, itemId)
	if err != nil {
		return fmt.Errorf("could not add item with id %d to player with user id %d: %w", itemId, userId, err)
	}
	return nil
}
