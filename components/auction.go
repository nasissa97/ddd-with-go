package components

import (
	"time"

	"github.com/Rhymond/go-money"
)

type Auction struct {
	ID            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

type Bid struct {
	ID        int
	offer     money.Money
	buyerID   int
	createdAt time.Time
	updatedAt time.Time
}

type User struct {
	ID        int
	name      string
	email     string // Email are unique in a system, however not a good idea to use them as Entity ID.
	createdAt time.Time
	updatedAt time.Time
}
