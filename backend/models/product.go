package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	ImageURL    string
	Price       float64   `gorm:"not null"`
	Type        string    `gorm:"type:enum('game_id','topup_card');not null"`
	Available   bool      `gorm:"default:true"`
	SellerID    uint      `gorm:"not null"`
	Seller      User      `gorm:"foreignKey:SellerID"`
	Orders      []Order   `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
