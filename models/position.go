package models

import (
	"stock/utils/error_utils"
	"time"
)

type Position struct {
	ID         uint      `json:"id" gorm:"primaryKey; autoIncrement"`
	StockId    uint      `json:"stock_id" gorm:"index"`
	Quantity   uint      `json:"quantity"`
	Price      float64   `json:"price"`
	Commission float64   `json:"commission"`
	OpenedAt   time.Time `json:"opened_at" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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

	return nil
}
