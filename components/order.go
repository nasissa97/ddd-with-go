package components

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type item struct {
	name string
}

type Order struct {
	items         []item
	taxAmount     money.Money
	discount      money.Money
	paymentCardID uuid.UUID
	customerID    uuid.UUID
	// marketingOptIn bool Adding marketingOptIn falls out of the Domain of an Order
	// This information can be used else where like User struct, but during a Order
	// transaction its best leave it out. You may include a seperate transaction to update
	// the marketingOptIn value for a user.
}
