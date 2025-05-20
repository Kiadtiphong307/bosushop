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

// ดึงสินค้าตาม ID
func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := database.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// สร้างสินค้าใหม่
func CreateProduct(product *models.Product) error {
	// ตรวจสอบชื่อซ้ำ
	var exist models.Product
	if err := database.DB.Where("name = ?", product.Name).First(&exist).Error; err == nil {
		return errors.New("ชื่อสินค้านี้มีอยู่แล้ว")
	}
	return database.DB.Create(product).Error
}

// แก้ไขสินค้า
func UpdateProduct(id uint, data *models.Product) (*models.Product, error) {
	product, err := GetProductByID(id)
	if err != nil {
		return nil, err
	}

	*product = *data // แทนค่าทั้งก้อน
	product.ID = id  // ไม่ให้ id หาย

	if err := database.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// ลบสินค้า
func DeleteProduct(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}
