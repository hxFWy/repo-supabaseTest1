package user

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"supabase-testProject1/internal/service/auth"
	"supabase-testProject1/internal/types"
)

type userService struct {
	db               *sql.DB
	userRepository   types.UserRepository
	playerRepository types.PlayerRepository
}

func NewUserService(
	db *sql.DB,
	userRepository types.UserRepository,
	playerRepository types.PlayerRepository,
) *userService {
	return &userService{
		db:               db,
		userRepository:   userRepository,
		playerRepository: playerRepository,
	}
}

func (s *userService) Login(ctx context.Context, loginPayload types.LoginUserPayload) (token string, statusCode int, err error) {

	fetchedUser, err := s.userRepository.GetUserByUsername(loginPayload.Username)
	if err != nil {
		return "", http.StatusBadRequest, fmt.Errorf("user with username %s not found", loginPayload.Username)
	}

	if !auth.ComparePasswords(fetchedUser.Password, []byte(loginPayload.Password)) {
		return "", http.StatusBadRequest, fmt.Errorf("user with username %s not found - invalid password", loginPayload.Username)
	}

	secret := []byte(os.Getenv("jwt_secret"))
	tokenString, err := auth.CreateJWT(secret, fetchedUser.Id)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("failed to create JWT token")
	}

	return tokenString, http.StatusOK, nil

}

func (s *userService) Register(ctx context.Context, registerPayload types.RegisterUserPayload) (statusCode int, err error) {

	existing, _ := s.userRepository.GetUserByUsername(registerPayload.Username)
	if existing != nil {
		return http.StatusBadRequest, fmt.Errorf("user with username %s already exists", registerPayload.Username)
	}

	hashedPassword, err := auth.HashPassword(registerPayload.Password)
	if err != nil {
		return http.StatusInternalServerError, err

	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// rollback helper
	rollback := func(err error) (int, error) {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	newUser, err := s.userRepository.CreateUserTx(ctx, tx, types.RegisterUserPayload{
		Username: registerPayload.Username,
		Password: hashedPassword,
		Position: registerPayload.Position,
	})
	if err != nil {
		return rollback(fmt.Errorf("failed to create user: %w", err))
	}

	err = s.playerRepository.CreatePlayerTx(ctx, tx, types.CreatePlayerPayload{
		User_id:  newUser.Id,
		Position: registerPayload.Position,
	})
	if err != nil {
		return rollback(fmt.Errorf("failed to create player: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(fmt.Errorf("failed to commit transaction: %w", err))
	}

	return http.StatusCreated, nil
}
