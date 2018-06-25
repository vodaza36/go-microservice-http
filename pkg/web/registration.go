package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vodaza36/go-microservice-http/pkg/domain"
	"github.com/vodaza36/go-microservice-http/pkg/service"
)

// RegistrationHandler class
type RegistrationHandler struct {
	userServicePort     int
	customerServicePort int
}

// NewRegistrationHandler constructor
func NewRegistrationHandler(userServicePort int, customerServicePort int) *RegistrationHandler {
	return &RegistrationHandler{userServicePort, customerServicePort}
}

// RegistrationHandlerPost handles the POST request
func (p *RegistrationHandler) RegistrationHandlerPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receive create registration request")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var registration domain.Registration
	err := decoder.Decode(&registration)

	if err != nil {
		log.Fatalf("Error parsing request: %v", err)
		http.Error(w, "Invalid request body!", 400)
	}

	log.Printf("Try to create Registration: %v", registration)
	RegistrationService := service.NewRegistrationService(p.userServicePort, p.customerServicePort)
	err = RegistrationService.Register(&registration)

	if err != nil {
		log.Fatalf("Error creating registration: %v", err)
		http.Error(w, "Invalid service call!", 503)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
