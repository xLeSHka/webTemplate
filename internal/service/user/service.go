package user

import (
	"backend/internal/interfaces"
)

type Service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *Service {
	return &Service{repository: repository}
}
