package shop

import (
	"database/sql"
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

func (r *Repository) GetItemsList() []*types.Item {
	newItem := &types.Item{
		Id:          1,
		Name:        "test",
		Slot:        "testslot",
		Cost:        100000,
		Skill_bonus: 1000,
	}

	items := make([]*types.Item, 1)
	items = append(items, newItem)

	return items
}
