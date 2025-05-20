package models

import "gorm.io/gorm"

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Product{},
		&Coupon{},
		&Order{},
	)
}
