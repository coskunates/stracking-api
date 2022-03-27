package services

import (
	"gorm.io/gorm"
	"stock/database"
	"stock/models"
	"stock/utils/error_utils"
)

func NewStockService() StockServiceInterface {
	return stockService{client: database.GetClient()}
}

type StockServiceInterface interface {
	List() (*[]models.Stock, *error_utils.RestErr)
}

type stockService struct {
	client *gorm.DB
}

func (s stockService) List() (*[]models.Stock, *error_utils.RestErr) {
	var stocks []models.Stock
	result := s.client.Find(&stocks).Scan(&stocks)
	if result.RowsAffected > 0 {
		return &stocks, nil
	} else {
		return nil, error_utils.NewNotFoundError("there is no stock", 18)
	}
}
