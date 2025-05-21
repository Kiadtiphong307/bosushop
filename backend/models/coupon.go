package models

import "time"

type Coupon struct {
	ID              uint      `gorm:"primaryKey"`
	Code            string    `gorm:"type:varchar(20);unique;not null"` 
	DiscountPercent int       `gorm:"not null"`
	MaxUsage        int       `gorm:"not null"`
	UsedCount       int       `gorm:"default:0"`
	ExpireAt        time.Time `gorm:"not null"`
	Orders          []Order   `gorm:"foreignKey:CouponCode;references:Code"`
	CreatedAt       time.Time
}
