package types

import "time"

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

type UserRepository interface {
	GetUserByUsername(username string) (*User, error)
	CreateUser(RegisterUserPayload) error
}
