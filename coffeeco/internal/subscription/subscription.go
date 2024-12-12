package subscription

import (
	coffeeco "coffeeco/internal"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Keeping it simple for now
type Subscription struct {
	ID                  uuid.UUID
	SubsriptionProducts []coffeeco.Product
}

func CreateSubsription(products []coffeeco.Product) (*Subscription, error) {
	if len(products) == 0 {
		return nil, fmt.Errorf("must have at least on product in the subscription")
	}
	subscriptionID := uuid.New()
	return &Subscription{
		ID:                  subscriptionID,
		SubsriptionProducts: products,
	}, nil
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetSubscriptionItems(ctx context.Context, subscriptionID uuid.UUID) ([]coffeeco.Product, error) {
	products, err := s.repo.GetSubscriptionItems(ctx, subscriptionID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s Service) InsertSubscription(ctx context.Context, subscription Subscription) error {
	if err := s.repo.InsertSubscription(ctx, subscription); err != nil {
		return nil
	}

	return nil
}
