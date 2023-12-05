package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     OrderAction
	Status        OrderStatus
	Transactions  []*Transaction
}

func NewOrder(orderId string, investor *Investor, asset *Asset, shares int, price float64, orderType OrderAction) *Order {

	return &Order{
		ID:            orderId,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        Open,
		Transactions:  []*Transaction{},
	}
}

func (o *Order) AddPendingShares(shares int) {
	o.PendingShares += shares
}

func (o *Order) CloseOrder() {
	o.Status = Closed
}

type OrderAction string

const (
	Buy  OrderAction = "BUY"
	Sell OrderAction = "SELL"
)

type OrderStatus string

const (
	Open     OrderStatus = "OPEN"
	Closed   OrderStatus = "CLOSED"
	Canceled OrderStatus = "CANCELED"
)
