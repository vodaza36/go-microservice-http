# go-microservice-http

A simple demo showing a Microservice architecture bases on synchronous HTTP communication.

## Starting the Microservices

```bash
cd $GOPATH/src/github.com/vodaza36/go-microservice-http
go run cmd/user/main.go -port 8090
go run cmd/customer/main.go -port 8091
go run cmd/registration/main.go -portRegistration 8092 -portUser 8090 -portCustomer 8091
```

## Test

Testing via Apache Bench (AB):

```bash
ab -n 100 -c 1 -p registration.post -T "application/json" http://localhost:8092/registration
```