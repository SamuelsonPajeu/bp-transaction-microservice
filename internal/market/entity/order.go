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
	Buy              OrderAction = "BUY"
	Sell             OrderAction = "SELL"
	UnknownOrderType OrderAction = "UNKNOWN_ORDER_TYPE"
)

func MapOrderAction(action string) OrderAction {
	actionMap := map[string]OrderAction{
		"BUY":  Buy,
		"SELL": Sell,
	}

	if val, found := actionMap[action]; found {
		return val
	}

	return UnknownOrderType
}

type OrderStatus string

const (
	Open               OrderStatus = "OPEN"
	Closed             OrderStatus = "CLOSED"
	Canceled           OrderStatus = "CANCELED"
	UnknownOrderStatus OrderStatus = "UNKNOWN_ORDER_STATUS"
)

func MapOrderStatus(action string) OrderStatus {
	actionMap := map[string]OrderStatus{
		"OPEN":     Open,
		"CLOSED":   Closed,
		"CANCELED": Canceled,
	}

	if val, found := actionMap[action]; found {
		return val
	}

	return UnknownOrderStatus
}
