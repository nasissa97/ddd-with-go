package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nasissa97/ddd-with-go/fundamentals"
	"github.com/nasissa97/ddd-with-go/fundamentals/oapi"
	"github.com/nasissa97/ddd-with-go/patterns"
)

const BaseURL string = "0.0.0.0:8080"

func failed_attempt() {
	server := fundamentals.NewServer()

	r := mux.NewRouter()

	h := oapi.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    BaseURL,
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	// provider := oapi.NewAuthProvider("example")
	localhost := oapi.WithBaseURL(BaseURL)
	// TODO: Figure out how to add PORT to Client. Otherwise Docker
	client, err := oapi.NewClientWithResponses(BaseURL, localhost)

	if err != nil {
		log.Printf("Failed to crate new client: %v", err)
	}

	ctx := context.Background()
	res, err := client.GetUsers(ctx)
	if err != nil {
		log.Printf("Failed to get users: %v", err)
	}

	fmt.Println(res)

}

func main() {
	myCar, err := patterns.BuildCar("tesla")
	if err != nil {
		log.Fatal(err)
	}
	myCar.Honk()

	_, err = patterns.BuildCar("Chevy")
	if err != nil {
		log.Print(err)
	}
}
