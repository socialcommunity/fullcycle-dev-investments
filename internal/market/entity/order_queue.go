package entity

type OrderQueue struct {
	Orders []*Order
}

func (orderQueue OrderQueue) Less(i, j int) bool {
	return orderQueue.Orders[i].Price < orderQueue.Orders[j].Price
}

func (orderQueue OrderQueue) Swap(i, j int) {
	orderQueue.Orders[i], orderQueue.Orders[j] = orderQueue.Orders[j], orderQueue.Orders[i]
}

func (orderQueue OrderQueue) Len() int {
	return len(orderQueue.Orders)
}

func (orderQueue *OrderQueue) Push(x interface{}) {
	orderQueue.Orders = append(orderQueue.Orders, x.(*Order))
}

func (orderQueue *OrderQueue) Pop() interface{} {
	old := orderQueue.Orders
	n := len(old)
	x := old[n-1]
	orderQueue.Orders = old[0 : n-1]
	return x
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{
		Orders: []*Order{},
	}
}
