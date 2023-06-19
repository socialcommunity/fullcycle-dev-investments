package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order            []*Order
	Transactions     []*Transaction
	OrdersChannel    chan *Order //input
	OrdersChannelOut chan *Order //output
	WaitGroup        *sync.WaitGroup
}

func NewBook(orderChannel chan *Order, orderChannelOut chan *Order, waitGroup *sync.WaitGroup) *Book {
	return &Book{
		Order:            []*Order{},
		Transactions:     []*Transaction{},
		OrdersChannel:    orderChannel,
		OrdersChannelOut: orderChannelOut,
		WaitGroup:        waitGroup,
	}
}

func (book *Book) Trade() {
	buyOrders := NewOrderQueue()
	sellOrders := NewOrderQueue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range book.OrdersChannel {
		if order.OrderType == "BUY" {

			heap.Push(buyOrders, order)

			if sellOrders.Len() > 0 && sellOrders.Orders[0].Price <= order.Price {
				sellOrder := sellOrders.Pop().(*Order)

				if sellOrder.PendingShares > 0 {

					transaction := NewTransaction(
						sellOrder,
						order,
						order.Shares,
						sellOrder.Price,
					)
					book.AddTransaction(transaction, book.WaitGroup)

					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					book.OrdersChannelOut <- sellOrder
					book.OrdersChannelOut <- order

					if sellOrder.PendingShares > 0 {
						sellOrders.Push(sellOrder)
					}
				}
			}
		} else if order.OrderType == "SELL" {

			sellOrders.Push(order)

			if buyOrders.Len() > 0 && buyOrders.Orders[0].Price >= order.Price {
				buyOrder := buyOrders.Pop().(*Order)

				if buyOrder.PendingShares > 0 {
					transaction := NewTransaction(
						order,
						buyOrder,
						buyOrder.Shares,
						buyOrder.Price,
					)

					book.AddTransaction(transaction, book.WaitGroup)
					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					book.OrdersChannelOut <- buyOrder
					book.OrdersChannelOut <- order

					if buyOrder.PendingShares > 0 {
						buyOrders.Push(buyOrder)
					}
				}
			}
		}
	}
}

func (book *Book) AddTransaction(transaction *Transaction, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	sellingShares := transaction.SellingOrder.PendingShares
	buyingShares := transaction.BuyingOrder.PendingShares

	minShares := sellingShares

	if buyingShares < sellingShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.ID, -minShares)
	transaction.SellingOrder.PendingShares -= minShares

	transaction.BuyingOrder.Investor.UpdateAssetPosition(transaction.BuyingOrder.Asset.ID, minShares)
	transaction.BuyingOrder.PendingShares -= minShares

	transaction.CalculateTotal()
	transaction.CloseBuyOrder()
	transaction.CloseSellOrder()

	book.Transactions = append(book.Transactions, transaction)
}
