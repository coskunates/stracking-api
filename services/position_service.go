package services

import (
	"gorm.io/gorm"
	"stock/database"
	"stock/models"
	"stock/utils/date_utils"
	"stock/utils/error_utils"
	"stock/utils/logger_utils"
)

func NewPositionService() PositionServiceInterface {
	return positionService{client: database.GetClient()}
}

type PositionServiceInterface interface {
	Get(models.Position) (*models.Position, *error_utils.RestErr)
	Create(models.Position) (*models.Position, *error_utils.RestErr)
	ClosePosition(models.ClosedPosition) (*models.ClosedPosition, *error_utils.RestErr)
}

type positionService struct {
	client *gorm.DB
}

func (p positionService) Get(position models.Position) (*models.Position, *error_utils.RestErr) {
	result := p.client.First(&position, "id = ?", position.ID)
	if result.RowsAffected > 0 {
		if err := result.Row().Scan(&position.ID,
			&position.StockId,
			&position.Quantity,
			&position.Price,
			&position.Commission,
			&position.OpenedAt,
			&position.CreatedAt,
			&position.UpdatedAt,
		); err != nil {
			logger_utils.Error("error when trying to save position", err)
			return nil, error_utils.NewInternalServerError("error when trying to save position", 6)
		}
	} else {
		return nil, error_utils.NewNotFoundError("position not found", 7)
	}

	return &position, nil
}

func (p positionService) Create(pr models.Position) (*models.Position, *error_utils.RestErr) {
	pr.CreatedAt = date_utils.GetNowAsString()
	pr.UpdatedAt = date_utils.GetNowAsString()

	if err := p.client.Save(&pr).Error; err != nil {
		logger_utils.Error("error when trying to save position", err)
		return nil, error_utils.NewInternalServerError("error when trying to save position", error_utils.DatabaseCreateError)
	}

	return &pr, nil
}

func (p positionService) ClosePosition(cp models.ClosedPosition) (*models.ClosedPosition, *error_utils.RestErr) {
	cp.ClosedAt = date_utils.GetNowAsString()
	cp.CreatedAt = date_utils.GetNowAsString()
	cp.UpdatedAt = date_utils.GetNowAsString()

	if err := p.client.Save(&cp).Error; err != nil {
		logger_utils.Error("error when trying to close position", err)
		return nil, error_utils.NewInternalServerError("error when trying to save position", error_utils.DatabaseCreateError)
	}

	if cp.Quantity > cp.SaleQuantity {
		p.client.Model(&models.Position{}).Where("id = ?", cp.PositionId).Updates(models.Position{
			Quantity: cp.Quantity - cp.SaleQuantity,
		})
	}

	return &cp, nil
}
