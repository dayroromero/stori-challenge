package models

import (
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	AccountID uint `gorm:"foreignKey:ID"`
	Amount    float64
	Date      time.Time
	Type      string
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}
