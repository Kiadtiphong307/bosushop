package models

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	ProductID  uint      `gorm:"not null"`
	Product    Product   `gorm:"foreignKey:ProductID"`
	CouponCode *string   `gorm:"type:varchar(20);index"` 
	Coupon     *Coupon   `gorm:"foreignKey:CouponCode;references:Code"`
	TotalPrice float64   `gorm:"not null"`
	Status     string    `gorm:"type:enum('pending','paid','canceled');default:'pending'"` 
	CreatedAt  time.Time
}
