package models

import (
	"fmt"
	"stock/utils/error_utils"
)

type ClosedPosition struct {
	ID             uint64  `json:"id" gorm:"primaryKey; autoIncrement"`
	PositionId     uint64  `json:"position_id" gorm:"index"`
	StockId        uint64  `json:"stock_id" gorm:"index"`
	Quantity       uint64  `json:"quantity"`
	Price          float64 `json:"price"`
	Commission     float64 `json:"commission"`
	SaleQuantity   uint64  `json:"sale_quantity"`
	SalePrice      float64 `json:"sale_price"`
	SaleCommission float64 `json:"sale_commission"`
	OpenedAt       string  `json:"opened_at" gorm:"index"`
	ClosedAt       string  `json:"closed_at"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func (p *ClosedPosition) Validate() *error_utils.RestErr {
	if p.PositionId <= 0 {
		return error_utils.NewBadRequestError("position id is required", 7)
	}

	if p.StockId <= 0 {
		return error_utils.NewBadRequestError("stock id is required", 8)
	}

	if p.Quantity <= 0 {
		return error_utils.NewBadRequestError("quantity is required", 9)
	}

	if p.Price <= 0 {
		return error_utils.NewBadRequestError("price is required", 10)
	}

	if p.Commission <= 0 {
		return error_utils.NewBadRequestError("commission is required", 11)
	}

	if p.SaleQuantity <= 0 {
		return error_utils.NewBadRequestError("sale quantity is required", 12)
	}

	if p.SalePrice <= 0 {
		return error_utils.NewBadRequestError("sale price is required", 13)
	}

	if p.SaleCommission <= 0 {
		return error_utils.NewBadRequestError("sale commission is required", 14)
	}

	if p.Quantity < p.SaleQuantity {
		return error_utils.NewBadRequestError(
			fmt.Sprintf("position quantity (%d) is less than sale quantity (%d)",
				p.Quantity,
				p.SaleQuantity,
			), 15)
	}

	return nil
}
