package services

import (
	"gorm.io/gorm"
	"stock/models"
	"stock/utils/error_utils"
	"stock/utils/logger_utils"
	"time"
)

func NewPositionService(client *gorm.DB) PositionServiceInterface {
	return positionService{client: client}
}

type PositionServiceInterface interface {
	Create(position models.Position) (*models.Position, *error_utils.RestErr)
}

type positionService struct {
	client *gorm.DB
}

func (p positionService) Create(pr models.Position) (*models.Position, *error_utils.RestErr) {
	pr.CreatedAt = time.Now().UTC()
	pr.OpenedAt = time.Now().UTC()
	pr.UpdatedAt = time.Now().UTC()

	if err := p.client.Save(&pr).Error; err != nil {
		logger_utils.Error("error when trying to save position", err)
		return nil, error_utils.NewInternalServerError("error when trying to save position", 6)
	}

	return &pr, nil
}
