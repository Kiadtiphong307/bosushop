package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	ImageURL    string
	Price       float64   `gorm:"not null"`
	Category    string    `gorm:"type:enum('game_id','topup_card');not null"`
	CategoryID  uint      `gorm:"not null"`
	Available   bool      `gorm:"default:true"`
	Stock       int       `gorm:"not null"`
	Orders      []Order   `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}