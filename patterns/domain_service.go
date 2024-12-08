package patterns

import "errors"

type Product struct {
	ID             int
	InStock        bool
	InSomeonesCart bool
}

func (p *Product) CanBeBought() bool {
	return p.InStock && !p.InSomeonesCart
}

type ShoppingCart struct {
	ID          int
	Products    []Product
	IsFull      bool
	MaxCartSize int
}

/*
Looks fine at first glance, however this Shopping couples to seperate entities together.
In this case the ShoppingCart implemented business logic while referencing Product.
Business logic can be implemented within entity only if references to other Domain Object
are within the same aggregate, otherwise implement the logic in the Domain Service. Use
Domain Service encapsulate complex behavoirs without overloading an single entity.
func (s *ShoppingCart) AddToCart(p Product) bool {
	if s.IsFull {
		return false
	}
	if p.CanBeBought() {
		s.Products = append(s.Products, p)
	}
	if s.MaxCartSize == len(s.Products) {
		s.IsFull = true
	}
	return true
}
*/

type CheckoutService struct {
	shoppingCart *ShoppingCart
}

func NewCheckoutService(shoppingCart *ShoppingCart) *CheckoutService {
	return &CheckoutService{shoppingCart: shoppingCart}
}

func (c CheckoutService) AddProductToBasket(p *Product) error {
	if c.shoppingCart.IsFull {
		return errors.New("cannot add to cart, cart is full")
	}
	if p.CanBeBought() {
		c.shoppingCart.Products = append(c.shoppingCart.Products, *p)
	}
	if c.shoppingCart.MaxCartSize == len(c.shoppingCart.Products) {
		c.shoppingCart.IsFull = true
	}

	return nil
}
