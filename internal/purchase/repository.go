package purchase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	coffeeco "github.com/slf-aobrien/coffeeco/internal"
	"github.com/slf-aobrien/coffeeco/internal/payment"
	"github.com/slf-aobrien/coffeeco/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Store(ctx context.Context, purchae Purchase) error
	Ping(ctx context.Context) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

type mongoPurchase struct {
	Id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	Total              money.Money
	PaymentMeans       payment.Means
	TimeOfPurchase     time.Time
	CardToken          *string
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")

	return &MongoRepository{
		purchases: purchases,
	}, nil
}

func (mr *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongoP := toMongoPurchase(purchase)
	_, err := mr.purchases.InsertOne(ctx, mongoP)
	if err != nil {
		return fmt.Errorf("failed to persist purchase: %w", err)
	}
	return nil
}

// this is to decouple our purchase aggregate from the Mongo Implementation
func toMongoPurchase(p Purchase) mongoPurchase {
	mongoP := mongoPurchase{
		Id:                 p.id,
		Store:              p.Store,
		ProductsToPurchase: p.ProductsToPurchase,
		Total:              p.total,
		PaymentMeans:       p.PaymentMeans,
		TimeOfPurchase:     p.timeOfPurchase,
		CardToken:          p.CardToken,
	}
	log.Println("\tpurchase.repository toMongoPurchase ")
	log.Println("\tid: ", mongoP.Id)
	log.Println("\tstore: ", mongoP.Store.ID)
	log.Println("\tProductsToPurchase: ", mongoP.ProductsToPurchase[0].ItemName)
	log.Println("\ttotal: ", mongoP.Total.Amount())
	log.Println("\tPayment Means: ", mongoP.PaymentMeans)
	log.Println("\tTime of Purchase: ", mongoP.TimeOfPurchase)
	log.Println("\tCard Token: ", mongoP.CardToken)
	return mongoP
}

func (mr *MongoRepository) Ping(ctx context.Context) error {
	if _, err := mr.purchases.EstimatedDocumentCount(ctx); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}
	return nil
}
