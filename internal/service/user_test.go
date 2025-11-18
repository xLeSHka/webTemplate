package service

import (
	"backend/internal/infra/queries"
	"backend/internal/mocks"
	"context"
	"errors"
	"testing"

	"github.com/alexedwards/argon2id"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mock := mocks.NewUserMock()
	type Test struct {
		Name  string
		User  queries.User
		Error error
	}
	service := NewUser(mock)
	passwordHash, err := argon2id.CreateHash("Very_strong_password1235", argon2id.DefaultParams)
	assert.Nil(t, err)
	mock.CreateUser(context.Background(), queries.User{ID: ulid.Make().String(), Email: "checkemail@gmail.com", PasswordHash: passwordHash})
	tests := []Test{
		Test{
			Name: "неуниклаьный email",
			User: queries.User{
				Email:        "checkemail@gmail.com",
				PasswordHash: "Very_strong_password1235",
			},
			Error: errors.New("такой email уже зарегистирован"),
		},
		Test{
			Name: "успешное создание",
			User: queries.User{
				Email:        "goodemail@gmail.com",
				PasswordHash: "Very_strong_password1235",
			},
			Error: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err = service.Create(context.Background(), test.User.Email, test.User.PasswordHash)
			assert.Equal(t, test.Error, err)
		})
	}
}
func TestGetByID(t *testing.T) {
	mock := mocks.NewUserMock()
	type Test struct {
		Name  string
		ID    string
		Error error
	}
	service := NewUser(mock)
	passwordHash, err := argon2id.CreateHash("Very_strong_password1235", argon2id.DefaultParams)
	assert.Nil(t, err)
	id1 := ulid.Make().String()
	mock.CreateUser(context.Background(), queries.User{ID: id1, Email: "checkemail@gmail.com", PasswordHash: passwordHash})
	tests := []Test{
		Test{
			Name:  "не существующий айди",
			ID:    ulid.Make().String(),
			Error: errors.New("user not found"),
		},
		Test{
			Name:  "успешное получение",
			ID:    id1,
			Error: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			user, err := service.GetByID(context.Background(), test.ID)
			assert.Equal(t, test.Error, err)
			if err == nil {
				assert.Equal(t, queries.User{ID: id1, Email: "checkemail@gmail.com", PasswordHash: passwordHash}, user)
			}
		})
	}
}
func TestGetByEmail(t *testing.T) {
	mock := mocks.NewUserMock()
	type Test struct {
		Name  string
		Email string
		Error error
	}
	service := NewUser(mock)
	passwordHash, err := argon2id.CreateHash("Very_strong_password1235", argon2id.DefaultParams)
	assert.Nil(t, err)
	id1 := ulid.Make().String()
	mock.CreateUser(context.Background(), queries.User{ID: id1, Email: "checkemail@gmail.com", PasswordHash: passwordHash})
	tests := []Test{
		Test{
			Name:  "не существующий email",
			Email: "bademail@gmail.com",
			Error: errors.New("user not found"),
		},
		Test{
			Name:  "успешное получение",
			Email: "checkemail@gmail.com",
			Error: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			user, err := service.GetByEmail(context.Background(), test.Email)
			assert.Equal(t, test.Error, err)
			if err == nil {
				assert.Equal(t, queries.User{ID: id1, Email: "checkemail@gmail.com", PasswordHash: passwordHash}, user)
			}
		})
	}
}
