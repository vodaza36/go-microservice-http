package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vodaza36/go-microservice-http/pkg/domain"
	"github.com/vodaza36/go-microservice-http/pkg/service"
)

// CustomerHandlerPost handles the POST request
func CustomerHandlerPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receive create customer request")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var customer domain.Customer
	err := decoder.Decode(&customer)

	if err != nil {
		log.Fatalf("Error parsing request: %v", err)
		http.Error(w, "Invalid request body!", 400)
	}

	log.Printf("Try to create customer: %v", customer)
	CustomerService := service.NewCustomerService()
	foundCustomer, err := CustomerService.CreateCustomer(&customer)

	if err != nil {
		log.Fatalf("Error creating customer: %s, error: %v", foundCustomer.ID, err)
		http.Error(w, "Invalid service call!", 503)
	}

	result, err := json.Marshal(&foundCustomer)

	if err != nil {
		log.Fatalf("Error parsing json response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
