package dto

import (
	"github.com/SamuelsonPajeu/bp-transaction-microservice/go/internal/market/entity"
)

type TradeInput struct {
	OrderID       string  `json:"order_id"`
	InvestorID    string  `json:"investor_id"`
	AssetId       string  `json:"asset_id"`
	CurrentShares int     `json:"current_shares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
	OrderTypeStr  string  `json:"order_type"`
	OrderType     entity.OrderAction
}

func (ti *TradeInput) SetOrderType() {
	ti.OrderType = entity.MapOrderAction(ti.OrderTypeStr)
}

type OrderOutput struct {
	OrderID           string `json:"order_id"`
	InvestorID        string `json:"investor_id"`
	AssetID           string `json:"asset_id"`
	OrderTypeStr      string `json:"order_type"`
	StatusStr         string `json:"status"`
	Partial           int    `json:"partial"`
	Shares            int    `json:"shares"`
	OrderType         entity.OrderAction
	Status            entity.OrderStatus
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

func (oo *OrderOutput) SetOrderType() {
	oo.OrderType = entity.MapOrderAction(oo.OrderTypeStr)
}

func (oo *OrderOutput) SetOrderStatus() {
	oo.Status = entity.MapOrderStatus(oo.StatusStr)
}

type TransactionOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerID       string  `json:"buyer_id"`
	SellerID      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}
