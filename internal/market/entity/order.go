package entity

type OrderType string

const (
	BUY  OrderType = "BUY"
	SELL OrderType = "SELL"
)

type OrderStatus string

const (
	OPEN    OrderStatus = "OPEN"
	PENDING OrderStatus = "PENDING"
	CLOSED  OrderStatus = "CLOSED"
)

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     OrderType
	Status        OrderStatus
	Transactions  []*Transaction
}

func NewOrder(
	id string,
	investor *Investor,
	asset *Asset,
	shares int,
	price float64,
	orderType OrderType,
) *Order {
	return &Order{
		ID:            id,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        OPEN,
		Transactions:  []*Transaction{},
	}
}
