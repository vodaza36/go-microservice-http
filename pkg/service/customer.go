package service

import (
	"log"

	"github.com/vodaza36/go-microservice-http/pkg/domain"
)

// CustomerService implementation
type CustomerService struct {
}

// NewCustomerService constructor
func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

// CreateCustomer service implementation
func (p *CustomerService) CreateCustomer(u *domain.Customer) (*domain.Customer, error) {
	log.Printf("Customer service successfully creates new customer: %v", *u)
	return &domain.Customer{
		ID:      u.ID,
		Name:    u.Name,
		Address: u.Address}, nil
}

// GetCustomerByID service implementation
func (p *CustomerService) GetCustomerByID(ID string) (*domain.Customer, error) {
	return &domain.Customer{}, nil
}
