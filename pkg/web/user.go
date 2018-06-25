package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vodaza36/go-microservice-http/pkg/domain"
	"github.com/vodaza36/go-microservice-http/pkg/service"
)

// UserHandlerPost handles the POST request
func UserHandlerPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receive create user request")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user domain.User
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatalf("Error parsing request: %v", err)
		http.Error(w, "Invalid request body!", 400)
	}

	log.Printf("Try to create user: %v", user)
	userService := service.NewUserService()
	foundUser, err := userService.CreateUser(&user)

	if err != nil {
		log.Fatalf("Error creating user: %s, error: %v", foundUser.ID, err)
		http.Error(w, "Invalid service call!", 503)
	}

	result, err := json.Marshal(&foundUser)

	if err != nil {
		log.Fatalf("Error parsing json response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
