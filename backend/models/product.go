package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	ImageURL    string 
	Slug        string    `gorm:"unique;not null"`
	Price       float64   `gorm:"not null"`
	CategoryID  uint      `gorm:"not null"`                          // FK
	Category    Category  `gorm:"foreignKey:CategoryID"`            // Join table
	Available   bool      `gorm:"default:true"`
	Stock       int       `gorm:"not null"`
	Orders      []Order   `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}