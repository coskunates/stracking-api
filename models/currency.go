package models

type Currency struct {
	ID        uint64 `json:"id" gorm:"primaryKey; autoIncrement"`
	Code      string `json:"code"`
	Symbol    string `json:"symbol"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
