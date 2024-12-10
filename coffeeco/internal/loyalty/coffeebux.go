package loyalty

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	coffeeco "coffeeco/internal"
	"coffeeco/internal/store"
)

type CoffeeBux struct {
	ID                                  uuid.UUID
	store                               store.Store
	coffeeLover                         coffeeco.CoffeeLover
	FreeDrinksAvailable                 int
	RemaingDrinkPurchasesUntilFreeDrink int
}

func (c *CoffeeBux) AddStamp() {
	if c.RemaingDrinkPurchasesUntilFreeDrink == 1 {
		c.RemaingDrinkPurchasesUntilFreeDrink = 1
		c.FreeDrinksAvailable += 1
	} else {
		c.RemaingDrinkPurchasesUntilFreeDrink--
	}
}

func (c *CoffeeBux) Pay(ctx context.Context, purchases []coffeeco.Product) error {
	lp := len(purchases)
	if lp == 0 {
		return errors.New("nothing to buy")
	}

	if c.FreeDrinksAvailable < lp {
		return fmt.Errorf("not enough coffeeBux to cover entire purchase. Have %d, need %d", len(purchases), c.FreeDrinksAvailable)
	}

	c.FreeDrinksAvailable = c.FreeDrinksAvailable - lp
	return nil
}
