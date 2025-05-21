package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/models"
)

var DB *gorm.DB

func InitDatabase() {
	// 🔃 โหลดไฟล์ .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ ไม่พบไฟล์ .env แต่จะดำเนินการต่อ...")
	}

	// ✅ ตรวจสอบว่าค่า env ไม่ว่างเปล่า
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		log.Fatal("❌ ไม่พบค่าที่จำเป็นใน .env: DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME")
	}

	// ✅ สร้าง DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	// ✅ เชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ ไม่สามารถเชื่อมต่อฐานข้อมูล: %v", err)
	}

	log.Println("✅ เชื่อมต่อฐานข้อมูลสำเร็จ")

	// 📦 เก็บไว้ในตัวแปร global
	DB = db

	// 🔃 เรียก Migration
	if err := models.MigrateTables(DB); err != nil {
		log.Fatalf("❌ Migration ล้มเหลว: %v", err)
	}
}
