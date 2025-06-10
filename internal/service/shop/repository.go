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
