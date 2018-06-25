package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/vodaza36/go-microservice-http/pkg/domain"
)

// RegistrationService class
type RegistrationService struct {
	userServicePort     int
	customerServicePort int
}

// NewRegistrationService constructor
func NewRegistrationService(userServicePort int, customerServicePort int) *RegistrationService {
	return &RegistrationService{userServicePort, customerServicePort}
}

// Register a new user
func (p *RegistrationService) Register(data *domain.Registration) error {
	newUser := domain.User{
		ID:    uuid.New().String(),
		Name:  data.Name,
		Email: data.Email,
	}
	err := p.invokeCreateUser(&newUser)

	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	newCustomer := domain.Customer{
		ID:      uuid.New().String(),
		Name:    data.Name,
		Address: data.Address,
	}

	err = p.invokeCreateCustomer(&newCustomer)

	if err != nil {
		log.Fatalf("Error creating customer: %v", err)
	}

	log.Printf("Successfully create user %s and customer %s", newUser.ID, newCustomer.ID)
	return err
}

func (p *RegistrationService) invokeCreateUser(user *domain.User) error {
	body, err := json.Marshal(&user)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%d/user", p.userServicePort), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error invoking the user service: %v", err)
	}
	defer response.Body.Close()
	return err
}

func (p *RegistrationService) invokeCreateCustomer(user *domain.Customer) error {
	body, err := json.Marshal(&user)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%d/customer", p.customerServicePort), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error invoking the customer service: %v", err)
	}
	defer response.Body.Close()
	return err
}
