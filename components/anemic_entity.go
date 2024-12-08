package components

import (
	"time"

	"github.com/Rhymond/go-money"
)

// Example of anemic model :/. Can be common in Go projects however you defeat the purpose of DDD.
// With anemic models what can happen is other models make assumption of the data and attempt to
// implement the business logic which can be in accuarate. Since communication between Bounded Context
// can occur, it's best not to let someone implement the business from a different Bounded Context.
type AnemicAuction struct {
	id            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

func (a *AnemicAuction) GetID() int {
	return a.id
}

func (a *AnemicAuction) StartingPrice() money.Money {
	return a.startingPrice
}

func (a *AnemicAuction) SetStartingPrice(startingPrice money.Money) {
	a.startingPrice = startingPrice
}

func (a *AnemicAuction) GetSellerID() int {
	return a.sellerID
}

func (a *AnemicAuction) SetSellerID(sellerID int) {
	a.sellerID = sellerID
}

func (a *AnemicAuction) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *AnemicAuction) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

func (a *AnemicAuction) GetAuctionStart() time.Time {
	return a.auctionStart
}

func (a *AnemicAuction) SetAuctionStart(auctionStart time.Time) {
	a.auctionStart = auctionStart
}

func (a *AnemicAuction) GetAuctionEnd() time.Time {
	return a.auctionEnd
}

func (a *AnemicAuction) SetAuctionEnd(auctionEnd time.Time) {
	a.auctionEnd = auctionEnd
}
