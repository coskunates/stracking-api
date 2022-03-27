package models

import "stock/utils/error_utils"

type Dividend struct {
	ID        uint64  `json:"id" gorm:"primaryKey; autoIncrement"`
	StockId   uint64  `json:"stock_id"`
	Quantity  uint64  `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func (d *Dividend) Validate() *error_utils.RestErr {
	if d.StockId <= 0 {
		return error_utils.NewBadRequestError("stock id is required", 19)
	}

	if d.Quantity <= 0 {
		return error_utils.NewBadRequestError("quantity is required", 20)
	}

	if d.Price <= 0 {
		return error_utils.NewBadRequestError("price is required", 21)
	}

	return nil
}
