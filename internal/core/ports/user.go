package ports

import "github.com/alandavd/empomio/internal/core/domain"

type UserPort interface {
	CreateUser() (*domain.User, error)
}
