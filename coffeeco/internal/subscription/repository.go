package subscription

import (
	coffeeco "coffeeco/internal"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetSubscriptionItems(ctx context.Context, subscriptionID uuid.UUID) ([]coffeeco.Product, error)
	InsertSubscription(ctx context.Context, subscription Subscription) error
}

type MongoRepository struct {
	subscriptions *mongo.Collection
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failded to create a mongo client: %v", err)
	}

	subscriptions := client.Database("coffeeco").Collection("subscriptions")
	return &MongoRepository{
		subscriptions: subscriptions,
	}, nil
}

func (mr MongoRepository) GetSubscriptionItems(
	ctx context.Context,
	subscriptionID uuid.UUID,
) ([]coffeeco.Product, error) {
	var subscriptionProducts = make([]coffeeco.Product, 0)
	var mSubscription mongoSubscription
	filter := bson.D{{Key: "ID", Value: subscriptionID.String()}}
	if err := mr.subscriptions.FindOne(ctx, filter).Decode(&mSubscription); err != nil {
		return subscriptionProducts, err
	}

	subscription := mSubscription.ToSubscription()

	return subscription.SubsriptionProducts, nil
}

func (mr MongoRepository) InsertSubscription(ctx context.Context, subscription Subscription) error {

	mongoSubsription := toMongoSubscription(subscription)
	_, err := mr.subscriptions.InsertOne(ctx, mongoSubsription)
	if err != nil {
		return fmt.Errorf("failed to create new service")
	}

	return nil
}

type mongoSubscription struct {
	ID                   uuid.UUID          `bson:"ID"`
	SubscriptionProducts []coffeeco.Product `bson:"subscription_products"`
}

func toMongoSubscription(s Subscription) mongoSubscription {
	return mongoSubscription{
		ID:                   s.ID,
		SubscriptionProducts: s.SubsriptionProducts,
	}
}

func (m mongoSubscription) ToSubscription() Subscription {
	return Subscription{
		ID:                  m.ID,
		SubsriptionProducts: m.SubscriptionProducts,
	}
}
