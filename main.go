// main.go
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
	"APP/"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", handleRequest).Methods("GET")
	r.HandleFunc("/api2", handleRequest2).Methods("GET")
	lambda.Start(handleRequest3)

	lambda.Start(r.ServeHTTP)
}
