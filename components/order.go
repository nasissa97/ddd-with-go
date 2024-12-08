package components

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type item struct {
	name string
}

type Order struct {
	items          []item
	taxAmount      money.Money
	discount       money.Money
	paymentCardID  uuid.UUID
	customerID     uuid.UUID
	marketingOptIn bool
}
