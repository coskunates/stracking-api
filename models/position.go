package models

import (
	"stock/utils/error_utils"
)

type Position struct {
	ID         uint64  `json:"id" gorm:"primaryKey; autoIncrement"`
	StockId    uint64  `json:"stock_id" gorm:"index"`
	Quantity   uint64  `json:"quantity"`
	Price      float64 `json:"price"`
	Commission float64 `json:"commission"`
	OpenedAt   string  `json:"opened_at" gorm:"index"`
	CreatedAt  string  `json:"created_at" gorm:"index"`
	UpdatedAt  string  `json:"updated_at" gorm:"index"`
}

func (p *Position) Validate() *error_utils.RestErr {
	if p.StockId <= 0 {
		return error_utils.NewBadRequestError("stock id is required", 2)
	}

	if p.Quantity <= 0 {
		return error_utils.NewBadRequestError("quantity is required", 3)
	}

	if p.Price <= 0 {
		return error_utils.NewBadRequestError("price is required", 4)
	}

	if p.Commission <= 0 {
		return error_utils.NewBadRequestError("commission is required", 5)
	}

	if p.OpenedAt == "" {
		return error_utils.NewBadRequestError("open time is required", 5)
	}

	return nil
}
