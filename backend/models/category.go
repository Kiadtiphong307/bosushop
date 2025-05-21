package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null;unique"`
	Products  []Product `gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time
}
