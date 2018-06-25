package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vodaza36/go-microservice-http/pkg/web"
)

var (
	portReg  = flag.Int("portRegistration", 8080, "Listen port of the registration service")
	portCust = flag.Int("portCustomer", 8080, "Listen port of the customer service")
	portUser = flag.Int("portUser", 8080, "Listen port of the user service")
)

func main() {
	flag.Parse()
	handler := web.NewRegistrationHandler(*portUser, *portCust)
	http.HandleFunc("/registration", handler.RegistrationHandlerPost)
	log.Printf("Listen to port: %d", *portReg)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portReg), nil))
}
