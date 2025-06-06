package types

import (
	"context"
	"database/sql"
	"time"
)

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Position string `json:"position"`
}

type LoginUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id         int       `json:"id"`
	Created_at time.Time `json:"created_at"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
}

type CreatePlayerPayload struct {
	User_id  int    `json:"user_id"`
	Position string `json:"position"`
}

type Player struct {
	User_id    int       `json:"user_id"`
	Money      float64   `json:"money"`
	Position   string    `json:"position"`
	Stamina    int       `json:"stamina"`
	Skill      int       `json:"skill"`
	Created_at time.Time `json:"created_at"`
}

type UserRepository interface {
	GetUserByUsername(username string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUserTx(ctx context.Context, tx *sql.Tx, payload RegisterUserPayload) (*User, error)
}

type PlayerRepository interface {
	GetPlayerByUserId(id int) (*Player, error)
	CreatePlayerTx(ctx context.Context, tx *sql.Tx, payload CreatePlayerPayload) error
}

type TrainingRepository interface {
	TrainPlayerById(user_id int) error
}
