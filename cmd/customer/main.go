package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vodaza36/go-microservice-http/pkg/web"
)

var port = flag.Int("port", 8080, "Listen port")

func main() {
	flag.Parse()
	http.HandleFunc("/customer", web.CustomerHandlerPost)
	log.Printf("Listen to port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
