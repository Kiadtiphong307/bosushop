package seed

import (
	"backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitSeeder(db *gorm.DB) {
	seedCategories(db)
	seedProducts(db)
	seedCoupons(db)
	seedAdmin(db)
}

// 🟦 1. เพิ่มหมวดหมู่สินค้า
func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Top-Up"},
		{Name: "Game ID"},
	}

	for _, c := range categories {
		var existing models.Category
		if err := db.Where("name = ?", c.Name).First(&existing).Error; err != nil {
			db.Create(&c)
			log.Println("✅ เพิ่มหมวดหมู่:", c.Name)
		}
	}
}

// 🟨 2. เพิ่มสินค้า (เชื่อมกับ CategoryID)
func seedProducts(db *gorm.DB) {
	products := []models.Product{
		{
			Name:        "Garena Topup 100",
			Description: "บัตรเติมเงิน Garena มูลค่า 100 บาท",
			ImageURL:    "https://cdn.example.com/garena-100.png",
			Price:       100,
			CategoryID:  1, // Top-Up
			Available:   true,
			Stock:       10,
		},
		{
			Name:        "Free Fire Game ID",
			Description: "ขายไอดีเกม Free Fire พร้อมใช้",
			ImageURL:    "https://cdn.example.com/freefire-id.png",
			Price:       350,
			CategoryID:  2, // Game ID
			Available:   true,
			Stock:       5,
		},
	}

	for _, p := range products {
		var existing models.Product
		if err := db.Where("name = ?", p.Name).First(&existing).Error; err != nil {
			db.Create(&p)
			log.Println("✅ เพิ่มสินค้า:", p.Name)
		}
	}
}

// 🟩 3. เพิ่มคูปอง
func seedCoupons(db *gorm.DB) {
	coupons := []models.Coupon{
		{
			Code:            "TOPUP10",
			DiscountPercent: 10,
			MaxUsage:        20,
			UsedCount:       0,
			ExpireAt:        time.Now().AddDate(0, 1, 0), // 1 เดือน
		},
		{
			Code:            "GAMEID20",
			DiscountPercent: 20,
			MaxUsage:        5,
			UsedCount:       0,
			ExpireAt:        time.Now().AddDate(0, 0, 15), // 15 วัน
		},
	}

	for _, c := range coupons {
		var existing models.Coupon
		if err := db.Where("code = ?", c.Code).First(&existing).Error; err != nil {
			db.Create(&c)
			log.Println("✅ เพิ่มคูปอง:", c.Code)
		}
	}
}

// 🟥 4. เพิ่มผู้ดูแลระบบ
func seedAdmin(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 12)
	admin := models.User{
		Username: "admin",
		Email:    "admin@bosushop.com",
		Password: string(password),
		Role:     "admin",
	}
	var exist models.User
	if err := db.Where("email = ?", admin.Email).First(&exist).Error; err != nil {
		db.Create(&admin)
		log.Println("✅ เพิ่มผู้ดูแลระบบ admin@bosushop.com (รหัสผ่าน: admin123)")
	}
}
