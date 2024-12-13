package main

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/purchase"
	"coffeeco/internal/store"
	"context"
	"log"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	// Stripe Documentation not real
	stripeTestAPIKey := "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	// Test token
	cardToken := "tok_visa"

	mongoConString := "mongodb://root:example@localhost:27017"
	csvc, err := payment.NewStripeService(stripeTestAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	purchaseRepo, err := purchase.NewMongRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := purchaseRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	storeRepo, err := store.NewMongoRep(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := storeRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	storeService := store.NewService(storeRepo)

	service := purchase.NewService(csvc, purchaseRepo, storeService)

	storeID := uuid.New()

	productPurchase := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: storeID,
		},
		ProductsToPurchase: []coffeeco.Product{{
			ItemName:  "item1",
			BasePrice: *money.New(3300, "USD"),
		}},
		PaymentMeans: payment.MEANS_CARD,
	}

	if err := service.CompletePurchase(ctx, storeID, productPurchase, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("product purchase was successful")

	subscriptionPurchase := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: storeID,
		},
		ProductsToPurchase: []coffeeco.Product{},
		SubscriptionsToPurchase: []coffeeco.Subscription{{
			SubscriptionTier: coffeeco.CoffeeDaily,
		}},
		PaymentMeans: payment.MEANS_CARD,
	}

	if err := service.CompletePurchase(ctx, storeID, subscriptionPurchase, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("subscription purchase was successful")

}
