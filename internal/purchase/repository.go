package purchase

import (
	"context"
	"fmt"
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
	id                 uuid.UUID
	store              store.Store
	productsToPurchase []coffeeco.Product
	total              money.Money
	paymentMeans       payment.Means
	timeOfPurchase     time.Time
	cardToken          *string
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
	return mongoPurchase{
		id:                 p.id,
		store:              p.Store,
		productsToPurchase: p.ProductsToPurchase,
		total:              p.total,
		paymentMeans:       p.PaymentMeans,
		timeOfPurchase:     p.timeOfPurchase,
		cardToken:          p.CardToken,
	}
}

func (mr *MongoRepository) Ping(ctx context.Context) error {
	if _, err := mr.purchases.EstimatedDocumentCount(ctx); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}
	return nil
}
