package user

import (
	"context"
	"supabase-testProject1/internal/types"
)

type UserService interface {
	Register(ctx context.Context, registerPayload types.RegisterUserPayload) (statusCode int, err error)

	Login(ctx context.Context, payload types.LoginUserPayload) (tokenString string, statusCode int, err error)
}
