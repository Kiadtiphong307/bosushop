package services

import (
	"backend/database"
	"backend/models"
	"errors"
)

// ดึงสินค้าทั้งหมด
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Order("created_at desc").Find(&products).Error
	return products, err
}

// สร้างสินค้าใหม่ (ใช้โดย controller)
func CreateProduct(product *models.Product) error {
	// ตรวจสอบชื่อซ้ำ
	var existing models.Product
	if err := database.DB.Where("name = ?", product.Name).First(&existing).Error; err == nil {
		return errors.New("ชื่อสินค้านี้มีอยู่แล้ว")
	}
	return database.DB.Create(product).Error
}
