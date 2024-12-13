package coffeeco

import (
	"errors"

	"github.com/Rhymond/go-money"
)

type SubscriptionTier string

const (
	CoffeeDaily   SubscriptionTier = "daily"
	CoffeeWeekly  SubscriptionTier = "weekly"
	CoffeeMonthly SubscriptionTier = "monthly"
)

func (tier SubscriptionTier) ToString() string {
	return string(tier)
}

func (tier SubscriptionTier) IsValid() bool {
	_, exists := AllowedSubscriptionTiers[tier]
	return exists
}

func ValidateSubscriptionTier(tier SubscriptionTier) error {
	if !tier.IsValid() {
		return errors.New("invalid subscription tier")
	}
	return nil
}

var AllowedSubscriptionTiers = map[SubscriptionTier]bool{
	CoffeeDaily:   true,
	CoffeeWeekly:  true,
	CoffeeMonthly: true,
}

type SubscriptionPrice *money.Money

var (
	DailyPrice   SubscriptionPrice = money.New(50, "USD")
	WeeklyPrice  SubscriptionPrice = money.New(25, "USD")
	MonthlyPrice SubscriptionPrice = money.New(10, "USD")
)

var subscriptionPriceMapping = map[SubscriptionTier]SubscriptionPrice{
	CoffeeDaily:   DailyPrice,
	CoffeeWeekly:  WeeklyPrice,
	CoffeeMonthly: MonthlyPrice,
}

// Keeping it simple for now
type Subscription struct {
	SubscriptionTier SubscriptionTier
}

func (subscription Subscription) SubscriptionCost() *money.Money {
	return subscriptionPriceMapping[subscription.SubscriptionTier]
}

/*
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
*/
