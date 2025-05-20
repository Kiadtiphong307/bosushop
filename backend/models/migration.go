package models

import (
	"log"

	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) error {
	log.Println("🔃 กำลังทำ Auto Migrate...")

	err := db.AutoMigrate(
		&User{},
		&Product{},
		&Coupon{},
		&Order{},
	)

	if err != nil {
		return err
	}

	log.Println("✅ Migration เสร็จสมบูรณ์")
	return nil
}
