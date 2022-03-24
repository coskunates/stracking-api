package services

import (
	"gorm.io/gorm"
	"stock/models"
	"time"
)

func NewStockService(client *gorm.DB) StockServiceInterface {
	return stockService{client: client}
}

type StockServiceInterface interface {
	Create(position models.Position) (*models.Position, error)
}

type stockService struct {
	client *gorm.DB
}

func (s stockService) Create(pr models.Position) (*models.Position, error) {
	pr.CreatedAt = time.Now().UTC()
	pr.OpenedAt = time.Now().UTC()
	pr.UpdatedAt = time.Now().UTC()

	if err := s.client.Save(&pr).Error; err != nil {
		return nil, err
	}

	return &pr, nil
}
