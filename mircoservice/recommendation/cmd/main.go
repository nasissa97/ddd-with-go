package main

import (
	"log"
	"net/http"

	"micro/recommendation/internal/recommendation"
	"micro/recommendation/internal/transport"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	c := retryablehttp.NewClient()
	c.RetryMax = 10

	// Required for Recommendation Service, to communicate with Partnerships MircoService
	partnerAdaptor, err := recommendation.NewPartnershipAdaptor(
		c.StandardClient(),
		"http://localhost:3031",
	)
	if err != nil {
		log.Fatal("faild to create a partnerAdaptor: ", err)
	}

	log.Println("Partnership Adaptor created")

	// Required for OpenHost Service to expose our Recommendation MircoService
	svc, err := recommendation.NewService(partnerAdaptor)
	if err != nil {
		log.Fatal("faild to create a service: ", err)
	}

	log.Println("Recommendation Service created")

	// Required to attach our Open Host Service to our router.
	handler, err := recommendation.NewHandler(*svc)
	if err != nil {
		log.Fatal("failed to create a handler: ", err)
	}

	log.Println("Open Host Service created")

	m := transport.NewMux(*handler)

	log.Println("Router Created")

	if err := http.ListenAndServe(":4040", m); err != nil {
		log.Fatal("server errored: ", err)
	}
}
