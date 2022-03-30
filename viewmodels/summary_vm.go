package viewmodels

import (
	"stock/models"
)

type SummaryVM struct {
	SubTotalPrice        float64      `json:"sub_total_price"`
	TotalPrice           float64      `json:"total_price"`
	TotalCommissionPrice float64      `json:"total_commission_price"`
	Items                []PositionVM `json:"items"`
}

type PositionVM struct {
	Id              uint64       `json:"id"`
	StockId         uint64       `json:"stock_id"`
	Stock           models.Stock `json:"stock"`
	Quantity        uint64       `json:"quantity"`
	Price           float64      `json:"price"`
	CommissionPrice float64      `json:"commission_price"`
	SubTotalPrice   float64      `json:"sub_total_price"`
	TotalPrice      float64      `json:"total_price"`
	CurrentPrice    float64      `json:"current_price"`
}

func (op *SummaryVM) BindPositionsToSummaryVM(positions *[]models.Position) *SummaryVM {
	var positionListVm SummaryVM

	for _, position := range *positions {
		positionListVm.SubTotalPrice += position.Price * float64(position.Quantity)
		positionListVm.TotalCommissionPrice += position.Commission

		positionVm := PositionVM{
			Id:              position.ID,
			StockId:         position.StockId,
			Stock:           position.Stock,
			Quantity:        position.Quantity,
			Price:           position.Price,
			CommissionPrice: position.Commission,
			SubTotalPrice:   position.Price * float64(position.Quantity),
			TotalPrice:      position.Price*float64(position.Quantity) + position.Commission,
		}
		positionListVm.Items = append(positionListVm.Items, positionVm)
	}

	positionListVm.TotalPrice = positionListVm.SubTotalPrice + positionListVm.TotalCommissionPrice

	return &positionListVm
}

func (op *SummaryVM) BindClosedPositionsToSummaryVM(positions *[]models.ClosedPosition) *SummaryVM {
	var positionListVm SummaryVM

	for _, position := range *positions {
		positionListVm.SubTotalPrice += position.Price * float64(position.Quantity)
		positionListVm.TotalCommissionPrice += position.Commission

		positionVm := PositionVM{
			Id:              position.ID,
			StockId:         position.StockId,
			Stock:           position.Stock,
			Quantity:        position.Quantity,
			Price:           position.Price,
			CommissionPrice: position.Commission,
			SubTotalPrice:   position.Price * float64(position.Quantity),
			TotalPrice:      position.Price*float64(position.Quantity) + position.Commission,
		}
		positionListVm.Items = append(positionListVm.Items, positionVm)
	}

	positionListVm.TotalPrice = positionListVm.SubTotalPrice + positionListVm.TotalCommissionPrice

	return &positionListVm
}
