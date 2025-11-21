package repository

import (
	"backend/internal/infra/queries"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user queries.User) error
	GetUserByID(ctx context.Context, id string) (queries.User, error)
	GetUserByEmail(ctx context.Context, email string) (queries.User, error)
}
