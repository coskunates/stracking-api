package services

import (
	"gorm.io/gorm"
	"stock/models"
	"stock/utils/error_utils"
	"stock/utils/logger_utils"
	"time"
)

func NewStockService(client *gorm.DB) StockServiceInterface {
	return stockService{client: client}
}

type StockServiceInterface interface {
	Create(position models.Position) (*models.Position, *error_utils.RestErr)
}

type stockService struct {
	client *gorm.DB
}

func (s stockService) Create(pr models.Position) (*models.Position, *error_utils.RestErr) {
	pr.CreatedAt = time.Now().UTC()
	pr.OpenedAt = time.Now().UTC()
	pr.UpdatedAt = time.Now().UTC()

	if err := s.client.Save(&pr).Error; err != nil {
		logger_utils.Error("error when trying to save position", err)
		return nil, error_utils.NewInternalServerError("error when trying to save position", 6)
	}

	return &pr, nil
}
