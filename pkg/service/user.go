package service

import (
	"log"

	"github.com/vodaza36/go-microservice-http/pkg/domain"
)

// UserService implementation
type UserService struct {
}

// NewUserService constructor
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser service implementation
func (p *UserService) CreateUser(u *domain.User) (*domain.User, error) {
	log.Printf("User service successfully creates new user: %v", *u)
	return &domain.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email}, nil
}

// GetUserByEmail service implementation
func (p *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return &domain.User{}, nil
}
