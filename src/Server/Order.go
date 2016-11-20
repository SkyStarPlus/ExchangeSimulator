package Server

import (
	"time"
)

const (
	Bid = iota
	Offer
)

type Order interface {
	Price() int	
	OrderType() int	
	Amount() int
	SetAmount(int) 
	ProductId() int
	OrderId() string
	SetOrderId(string)
	User() User
	Time() time.Time
}

type OrderImpl struct {
	_price int
	_orderType int
	_amount int
	_productId int
	_orderId string
	_user User
	_time time.Time
}

func (s *OrderImpl) Price() int {
	return s._price
}

func (s *OrderImpl) OrderType() int {
	return s._orderType
}

func (s *OrderImpl) Amount() int {
	return s._amount
}

func (s *OrderImpl) SetAmount(amount int) {
	s._amount = amount
}

func (s *OrderImpl) ProductId() int {
	return s._productId
}

func (s *OrderImpl) OrderId() string {
	return s._orderId
}

func (s *OrderImpl) SetOrderId(id string) {
	s._orderId = id
}

func (s *OrderImpl) User() User{
	return s._user
}

func (s *OrderImpl) Time() time.Time {
	return s._time
}
func CreateOrder(price int, orderType int, amount int, productId int, user User) Order {
	order := new(OrderImpl)
	order._price = price
	order._orderType = orderType
	order._amount = amount
	order._productId = productId
	order._user = user
	order._time = time.Now()
	order._orderId = time.Now().String()
	return order
}