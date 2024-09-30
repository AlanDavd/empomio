package services

import (
	"github.com/alandavd/empomio/internal/core/domain"
	"github.com/alandavd/empomio/internal/core/ports"
)

type UserService struct {
	userRepository ports.UserPort
}

func NewUserService(repo ports.UserPort) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (s *UserService) CreateUser() (*domain.User, error) {
	return s.userRepository.CreateUser()
}