package transformer

import (
	"github.com/SamuelsonPajeu/bp-transaction-microservice/go/internal/market/dto"
	"github.com/SamuelsonPajeu/bp-transaction-microservice/go/internal/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {
	input.SetOrderType()

	asset := entity.NewAsset(input.AssetId, input.AssetId, 1000)
	investor := entity.NewInvestor(input.InvestorID)
	order := entity.NewOrder(
		input.OrderID,
		investor,
		asset,
		input.Shares,
		input.Price,
		input.OrderType,
	)

	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetId, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}

	return order
}

func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderID:      order.ID,
		InvestorID:   order.Investor.ID,
		AssetID:      order.Asset.ID,
		OrderType:    order.OrderType,
		OrderTypeStr: string(order.OrderType),
		Status:       order.Status,
		StatusStr:    string(order.Status),
		Partial:      order.PendingShares,
		Shares:       order.Shares,
	}

	var transactionsOutput []*dto.TransactionOutput
	for _, t := range order.Transactions {
		transactionOutput := &dto.TransactionOutput{
			TransactionID: t.ID,
			BuyerID:       t.BuyingOrder.ID,
			SellerID:      t.SellingOrder.ID,
			AssetID:       t.SellingOrder.ID,
			Price:         t.Price,
			Shares:        t.SellingOrder.Shares - t.SellingOrder.PendingShares,
		}
		transactionsOutput = append(transactionsOutput, transactionOutput)
	}

	output.TransactionOutput = transactionsOutput
	return output
}
