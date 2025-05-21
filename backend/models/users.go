package models

import "time"

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"unique;not null"`
	Email     string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Role      string  `gorm:"type:enum('user','admin');default:'user';not null"`
	Orders    []Order `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}