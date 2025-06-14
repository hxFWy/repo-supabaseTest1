package shop

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"supabase-testProject1/internal/types"
)

type userService struct {
	db               *sql.DB
	userRepository   types.UserRepository
	playerRepository types.PlayerRepository
	shopRepository   types.ShopRepository
}

func NewService(
	db *sql.DB,
	userRepository types.UserRepository,
	playerRepository types.PlayerRepository,
	shopRepository types.ShopRepository,
) *userService {
	return &userService{
		db:               db,
		userRepository:   userRepository,
		playerRepository: playerRepository,
		shopRepository:   shopRepository,
	}
}

func (s *userService) PurchaseItem(ctx context.Context, userId, itemId int) (statusCode int, err error) {

	item, err := s.shopRepository.GetItemById(ctx, itemId)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("could not retrieve item with id %d: %w", itemId, err)
	}

	player, err := s.playerRepository.GetPlayerByUserId(userId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not retrieve player for user id %d: %w", userId, err)
	}

	if player.Money < float64(item.Cost) {
		return http.StatusForbidden, fmt.Errorf("not enough money to purchase item with id %d", itemId)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := s.playerRepository.UpdatePlayerMoneyTx(ctx, tx, userId, float64(-item.Cost)); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not update player money: %w", err)
	}

	if err := s.shopRepository.AddItemToPlayerTx(ctx, tx, userId, itemId); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not add item to player: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not commit transaction: %w", err)
	}

	return http.StatusCreated, nil

}
