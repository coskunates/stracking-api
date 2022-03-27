package models

import "stock/utils/error_utils"

type Stock struct {
	ID        uint64 `json:"id" gorm:"primaryKey; autoIncrement"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Country   string `json:"country"`
	Exchange  string `json:"exchange"`
	Currency  string `json:"currency"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *Stock) Validate() *error_utils.RestErr {
	if s.Name == "" {
		return error_utils.NewBadRequestError("stock name is required", 15)
	}

	if s.ShortName == "" {
		return error_utils.NewBadRequestError("stock short name is required", 16)
	}

	if s.Country == "" {
		return error_utils.NewBadRequestError("stock country is required", 17)
	}

	if s.Currency == "" {
		return error_utils.NewBadRequestError("currency is required", 19)
	}

	return nil
}
