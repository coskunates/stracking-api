package services

import (
	"gorm.io/gorm"
	"stock/database"
	"stock/models"
	"stock/utils/error_utils"
)

func NewSummaryService() SummaryServiceInterface {
	return summaryService{client: database.GetClient()}
}

type SummaryServiceInterface interface {
	Portfolio() (*[]models.Position, *error_utils.RestErr)
	OpenPositions() (*[]models.Position, *error_utils.RestErr)
	ClosedPositions() (*[]models.ClosedPosition, *error_utils.RestErr)
}

type summaryService struct {
	client *gorm.DB
}

func (s summaryService) OpenPositions() (*[]models.Position, *error_utils.RestErr) {
	var positions []models.Position

	result := s.client.Joins("Stock").Order("opened_at desc").Find(&positions).Scan(&positions)
	if result.RowsAffected > 0 {
		return &positions, nil
	} else {
		return nil, error_utils.NewNotFoundError("there is no open position", 18)
	}
}

func (s summaryService) ClosedPositions() (*[]models.ClosedPosition, *error_utils.RestErr) {
	var positions []models.ClosedPosition

	result := s.client.Joins("Stock").Order("closed_at desc").Find(&positions).Scan(&positions)
	if result.RowsAffected > 0 {
		return &positions, nil
	} else {
		return nil, error_utils.NewNotFoundError("there is no closed position", 18)
	}
}

func (s summaryService) Portfolio() (*[]models.Position, *error_utils.RestErr) {
	var positions []models.Position

	result := s.client.Select("stock_id," +
		"sum(positions.price*quantity)/sum(quantity) as price," +
		"sum(quantity) as quantity," +
		"sum(commission) as commission").
		Group("stock_id").
		Joins("Stock").
		Find(&positions).
		Scan(&positions)

	if result.RowsAffected > 0 {
		return &positions, nil
	} else {
		return nil, error_utils.NewNotFoundError("no summary information", 18)
	}
}
