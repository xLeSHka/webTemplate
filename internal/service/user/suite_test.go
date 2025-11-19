package user

import (
	interfaceMocks "backend/internal/interfaces/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	userRepository *interfaceMocks.MockUserRepository
	service        *Service
}

func (s *ServiceSuite) SetupTest() {
	s.userRepository = interfaceMocks.NewMockUserRepository(s.T())
	s.service = NewService(s.userRepository)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
