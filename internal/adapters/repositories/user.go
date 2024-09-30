package repositories

import (
	"github.com/alandavd/empomio/internal/core/domain"
)

type userRepo struct {}

func NewUserRepository() (*userRepo, error) {
	return &userRepo{}, nil
}

func (r *userRepo) CreateUser() (*domain.User, error) {
	user := &domain.User{}
	// Connect to DB and make queries here
	return user, nil
}
