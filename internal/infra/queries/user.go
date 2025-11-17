package queries

import (
	"context"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) CreateUser(ctx context.Context, user User) error {
	err := r.db.WithContext(ctx).Create(&user).Error
	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("LOWER(users.email) = ?", strings.ToLower(email)).First(&user).Error
	return user, err
}
func (r *UserRepository) GetUserByID(ctx context.Context, id string) (User, error) {
	var user User
	err := r.db.WithContext(ctx).First(&user, &User{ID: id}).Error
	return user, err
}
