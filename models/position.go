package models

import (
	"github.com/labstack/echo/v4"
	"net/http"
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

func (p *Position) Validate() error {
	if p.StockId <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "stock id is required")
	}

	if p.Quantity <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "quantity is required")
	}

	if p.Price <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "price is required")
	}

	if p.Commission <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "commission is required")
	}
	return nil
}
