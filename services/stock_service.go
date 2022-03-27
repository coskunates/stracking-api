package services

import (
	"gorm.io/gorm"
	"stock/database"
	"stock/models"
	"stock/utils/date_utils"
	"stock/utils/error_utils"
	"stock/utils/logger_utils"
)

func NewStockService() StockServiceInterface {
	return stockService{client: database.GetClient()}
}

type StockServiceInterface interface {
	List() (*[]models.Stock, *error_utils.RestErr)
	Create(stock models.Stock) (*models.Stock, *error_utils.RestErr)
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

func (s stockService) Create(stock models.Stock) (*models.Stock, *error_utils.RestErr) {
	stock.CreatedAt = date_utils.GetNowAsString()
	stock.UpdatedAt = date_utils.GetNowAsString()

	if err := s.client.Save(&stock).Error; err != nil {
		logger_utils.Error("error when trying to save stock", err)
		return nil, error_utils.NewInternalServerError("error when trying to save stock", error_utils.DatabaseCreateError)
	}

	return &stock, nil
}
