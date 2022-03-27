package services

import (
	"gorm.io/gorm"
	"stock/database"
	"stock/models"
	"stock/utils/date_utils"
	"stock/utils/error_utils"
	"stock/utils/logger_utils"
)

func NewDividendService() DividendServiceInterface {
	return dividendService{client: database.GetClient()}
}

type DividendServiceInterface interface {
	Create(models.Dividend) (*models.Dividend, *error_utils.RestErr)
}

type dividendService struct {
	client *gorm.DB
}

func (d dividendService) Create(dividend models.Dividend) (*models.Dividend, *error_utils.RestErr) {
	dividend.CreatedAt = date_utils.GetNowAsString()
	dividend.UpdatedAt = date_utils.GetNowAsString()

	if err := d.client.Save(&dividend).Error; err != nil {
		logger_utils.Error("error when trying to save dividend", err)
		return nil, error_utils.NewInternalServerError("error when trying to save dividend", error_utils.DatabaseCreateError)
	}

	return &dividend, nil
}
