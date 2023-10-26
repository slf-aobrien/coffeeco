package main

import (
	"context"
	"log"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"

	coffeeco "github.com/slf-aobrien/coffeeco/internal"
	"github.com/slf-aobrien/coffeeco/internal/payment"
	"github.com/slf-aobrien/coffeeco/internal/purchase"
	"github.com/slf-aobrien/coffeeco/internal/store"
)

func main() {
	log.Println("Starting application")
	ctx := context.Background()
	log.Println("Context created")
	// This is the test key from Stripe's documentation. Feel free to use it, no charges will actually be made.
	stripeTestAPIKey := "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	// This is a test token from Stripe's documentation. Feel free to use it, no charges will actually be made.
	cardToken := "tok_visa"

	// This is the credentials for mongo if you run docker-compose up in this repo.
	log.Println("Creating Repos..")
	mongoConString := "mongodb://root:example@localhost:27017"
	csvc, err := payment.NewStripeService(stripeTestAPIKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created New Stripe Service")
	prepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Purchase Mongo Service")
	if err := prepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Created Store Mongo Service Ping")
	sRepo, err := store.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Store Mongo Service")
	if err := sRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Created Service Ping")

	sSvc := store.NewService(sRepo)
	log.Println("Created store service")
	svc := purchase.NewService(csvc, prepo, sSvc)
	log.Println("Created Purchae Service")
	someStoreID := uuid.New()
	log.Println("Started Purchase")
	pur := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: someStoreID,
		},
		ProductsToPurchase: []coffeeco.Product{{
			ItemName:  "item1",
			BasePrice: *money.New(3300, "USD"),
		}},
		PaymentMeans: payment.MEANS_CARD,
	}
	log.Println("Purchase data created")
	log.Println("Saving data..")
	if err := svc.CompletePurchase(ctx, someStoreID, pur, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Purchase saved")

	log.Println("purchase was successful")
}
