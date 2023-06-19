package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(
	sellingOrder *Order,
	buyingOrder *Order,
	shares int,
	price float64) *Transaction {

	total := float64(shares) * price

	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (transaction *Transaction) CalculateTotal() {
	transaction.Total = float64(transaction.Shares) * transaction.Price
}

func (transaction *Transaction) CloseBuyOrder() {

	if transaction.BuyingOrder.PendingShares == 0 {
		transaction.BuyingOrder.Status = "CLOSED"
	}
}
func (transaction *Transaction) CloseSellOrder() {

	if transaction.SellingOrder.PendingShares == 0 {
		transaction.SellingOrder.Status = "CLOSED"
	}
}

func (transaction *Transaction) AddBuyOrderPendingShares(shares int) {
	transaction.BuyingOrder.PendingShares += shares
}

func (transaction *Transaction) AddSellOrderPendingShares(shares int) {
	transaction.SellingOrder.PendingShares += shares
}
