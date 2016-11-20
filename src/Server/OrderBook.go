package Server

import (
	"sort"
)

type OrderBook interface {
	AddOrder(order Order)
	DelOrder(order Order)
	//ModOrderAmount(order Order, amount int)
	FindOrder(orderId string) (bool, Order)
	
	//AllOrders() map[string]Order
	BidOrders() []Order
	OfferOrders() []Order
	
	ProductId() int
	SetProductId(int)
	
	BestBidPrice() int
	BestOfferPrice() int
	
	BestBidOrder() Order
	BestOfferOrder() Order
	
	//Init()
}

type OrderBookImpl struct {
	//TODO
	_productId int
	_orderMap map[string]Order
	_bidOrderSeq []Order
	_offerOrderSeq []Order
}

func compareBidOrder(order1, order2 Order) bool {
	if (order1.Price() > order2.Price()) {
		return true
	} else if (order1.Price() == order2.Price()) {
		if (order1.Time().Before(order2.Time())) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func compareOfferOrder(order1, order2 Order) bool {
	if (order1.Price() < order2.Price()) {
		return true
	} else if (order1.Price() == order2.Price()) {
		if (order1.Time().Before(order2.Time())) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (o *OrderBookImpl) AddOrder(order Order) {
	//TODO 将order数据存入数据库中的orderBook 数据库操作失败是要抛异常的
	o._orderMap[order.OrderId()] = order
	if (order.OrderType() == Bid) {
		o._bidOrderSeq = append(o._bidOrderSeq, order)
		sort.Sort(OrderSorter{o._bidOrderSeq, compareBidOrder})
	} else {
		o._offerOrderSeq = append(o._offerOrderSeq, order)
		sort.Sort(OrderSorter{o._offerOrderSeq, compareOfferOrder})
	}
}

func (o *OrderBookImpl) DelOrder(order Order) {
	//TODO 将order数据从数据库中的orderBook删除 数据库操作失败是要抛异常的
	_, exist := o._orderMap[order.OrderId()]
	if (exist == false) {
		return 
	}
	delete(o._orderMap, order.OrderId())
	//从序列中将Order删除
	if (order.OrderType() == Bid) {
		var index = 0
		for ; index < len(o._bidOrderSeq); index++ {
			if (o._bidOrderSeq[index].OrderId() == order.OrderId()) {
				break
			}
		}
		o._bidOrderSeq = append(o._bidOrderSeq[:index], o._bidOrderSeq[index + 1:]...)
	} else {
		var index = 0
		for ; index < len(o._offerOrderSeq); index++ {
			if (o._offerOrderSeq[index].OrderId() == order.OrderId()) {
				break
			}
		}
		o._offerOrderSeq = append(o._offerOrderSeq[:index], o._offerOrderSeq[index + 1:]...)
	}
}

func (o *OrderBookImpl) ModOrderAmount(order Order, amount int) {
	if (order.OrderType() == Bid) {
		//TODO 在数据库中将数量数据改掉
		var index = 0
		for ; index < len(o._bidOrderSeq); index++ {
			if (o._bidOrderSeq[index].OrderId() == order.OrderId()) {
				break
			}
		}
		o._bidOrderSeq[index].SetAmount(amount)
	} else {
		//TODO 在数据库中将数量数据改掉
		var index = 0
		for ; index < len(o._offerOrderSeq); index++ {
			if (o._offerOrderSeq[index].OrderId() == order.OrderId()) {
				break
			}
		}
		o._offerOrderSeq[index].SetAmount(amount)
	}
}

func (o *OrderBookImpl) FindOrder(orderId string) (bool, Order) {
	order, exist := o._orderMap[orderId]
	if (exist) {
		return true, order
	} else {
		//TODO不存在应该是要 抛异常的
		return false, order
	}
}

func (o *OrderBookImpl) AllOrders() map[string]Order {
	return o._orderMap
}

func (o *OrderBookImpl) BidOrders() []Order {
	return o._bidOrderSeq
}

func (o *OrderBookImpl) OfferOrders() []Order {
	return o._offerOrderSeq
}

func (o *OrderBookImpl) ProductId() int {
	return o._productId
}

func (o *OrderBookImpl) SetProductId(productId int) {
	o._productId = productId
}

func (o *OrderBookImpl) BestBidPrice() int {
	if (len(o._bidOrderSeq) > 0) {
		return o._bidOrderSeq[0].Price()
	} else {
		return 0
	}
}

func (o *OrderBookImpl) BestOfferPrice() int {
	if (len(o._offerOrderSeq) > 0) {
		return o._offerOrderSeq[0].Price()
	} else {
		return 1000000
	}
}

//这里有问题，没有任何Order的情况下，应该构造一个非法的或者报错
func (o *OrderBookImpl) BestBidOrder() Order {
	if (len(o._bidOrderSeq) > 0) {
		return o._bidOrderSeq[0]
	} else {
		return o._bidOrderSeq[0]
	}
}

//这里有问题，没有任何Order的情况下，应该构造一个非法的或者报错
func (o *OrderBookImpl) BestOfferOrder() Order {
	if (len(o._offerOrderSeq) > 0) {
		return o._offerOrderSeq[0]
	} else {
		return o._offerOrderSeq[0]
	}
}

func (o *OrderBookImpl) Init() {
	o._orderMap = make(map[string]Order)
	o._bidOrderSeq = make([]Order, 0)//TODO 正确性有待检验
	o._offerOrderSeq = make([]Order, 0)//TODO 正确性有待检验
	//TODO 从数据库加载数据
	//对bidOrderSeq和offerOrderSeq进行排序 排序标准是价格和时间 
	sort.Sort(OrderSorter{o._bidOrderSeq, compareBidOrder})
	sort.Sort(OrderSorter{o._offerOrderSeq, compareOfferOrder})
	
}

func CreateOrderBook(productId int) OrderBook{
	orderBook := new(OrderBookImpl)
	orderBook.SetProductId(productId)
	orderBook.Init()
	return orderBook
}

/*
func (o *OrderBookImpl) SortOrder(compare func (order1, order2 *Order) bool) {
	o._orderSeq = make([]Order, len(o._orderMap))
	i := 0
	for _, value := range o._orderMap {
		o._orderSeq[i] = value
		i++	
	}
	sort.Sort(OrderSorter{o._orderSeq, compare})
}
*/

type OrderSorter struct {
	_orders []Order
	_compare func(order1, order2 Order) bool
}

func (orderSorter OrderSorter) Len() int {
	return len(orderSorter._orders)
}

func (orderSorter OrderSorter) Swap(i, j int) {
	orderSorter._orders[i], orderSorter._orders[j] = orderSorter._orders[j], orderSorter._orders[i]
}

func (orderSorter OrderSorter) Less(i, j int) bool {
	return orderSorter._compare(orderSorter._orders[i], orderSorter._orders[j])//这里可能有问题  可能要用指针
}

//OrderSorter排序的使用方法，sort.Sort( OrderSorter{orders, func (order1, order2 *Order) bool { 这里是对两个order的比较} } )