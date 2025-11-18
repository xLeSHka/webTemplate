package mocks

import (
	"backend/internal/infra/queries"
	"context"
	"errors"
)

type UserMock struct {
	Users map[string]queries.User
}

func NewUserMock() *UserMock {
	return &UserMock{
		Users: make(map[string]queries.User),
	}
}
func (um *UserMock) CreateUser(ctx context.Context, user queries.User) error {
	if _, ok := um.Users[user.ID]; ok {
		return errors.New("id должен быть уникальным")
	}
	for _, u := range um.Users {
		if u.Email == user.Email {
			return errors.New("такой email уже зарегистирован")
		}
	}
	um.Users[user.ID] = user
	return nil
}
func (um *UserMock) GetUserByID(ctx context.Context, id string) (queries.User, error) {
	u, ok := um.Users[id]
	if !ok {
		return queries.User{}, errors.New("user not found")
	}
	return u, nil
}
func (um *UserMock) GetUserByEmail(ctx context.Context, email string) (queries.User, error) {
	for _, u := range um.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return queries.User{}, errors.New("user not found")
}
