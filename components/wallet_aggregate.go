package components

import (
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type WalletItem interface {
	GetBalance() (money.Money, error)
}

type Wallet struct {
	// This is our aggregate root, and is our wallet's identify
	id uuid.UUID
	// Owner ID is the identitfy of the entity who owns the wallet. We do not need to all all the datail
	// of an owner at all time but this gives us the ability to get hem when nessary(Credit Card Denied).
	ownerID uuid.UUID
	// Wallet Item is an entity we defined elsewhere, since things such as its balance can change over time.
	walletItems []WalletItem
}

func (w Wallet) GetWalletBalance() (*money.Money, error) {
	var bal *money.Money
	for _, v := range w.walletItems {
		itemBal, err := v.GetBalance()
		if err != nil {
			return nil, errors.New("failed to get balance")
		}
		bal, err = bal.Add(&itemBal)
		if err != nil {
			return nil, errors.New("failed to increment balance")
		}
	}
	return bal, nil
}
