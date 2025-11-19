package interfaces

import (
	"backend/internal/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
}

type AuthService interface {
	VerifyToken(authHeader string) (string, error)
	VerifyPassword(user model.User, password string) error
	GenerateToken(userID string) (string, error)
}

type UserService interface {
	Create(ctx context.Context, email, password string) (string, error)
	GetByID(ctx context.Context, id string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}
