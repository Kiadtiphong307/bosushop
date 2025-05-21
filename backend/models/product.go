package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	ImageURL    string    `gorm:"column:image_url" json:"image_url"`
	Slug 		string 	  `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Price       float64   `gorm:"not null" json:"price"`
	CategoryID  uint      `gorm:"not null" json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Available   bool      `gorm:"default:true" json:"available"`
	Stock       int       `gorm:"not null" json:"stock"`
	Orders      []Order   `gorm:"foreignKey:ProductID" json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
