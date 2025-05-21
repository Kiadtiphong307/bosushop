package services

import (
	"backend/database"
	"backend/models"
	"errors"
)

// ดึงสินค้าทั้งหมด
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	Preload("Category"). 
	err := database.DB.Order("created_at desc").Find(&products).Error
	return products, err
}

// ดึงสินค้าตาม slug
func GetProductBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var product models.Product
	if err := database.DB.
		Preload("Category").
		Where("slug = ?", slug).
		First(&product).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "ไม่พบสินค้า"})
	}
	return c.JSON(product)
}

// สร้างสินค้าใหม่
func CreateProduct(product *models.Product) error {
	// ✅ สร้าง slug ภาษาไทย
	product.Slug = slug.MakeLang(product.Name, "th")

	// ตรวจสอบ slug ซ้ำ
	var existSlug models.Product
	if err := database.DB.Where("slug = ?", product.Slug).First(&existSlug).Error; err == nil {
		return errors.New("มีสินค้าที่ใช้ slug เดียวกันแล้ว")
	}

	// ตรวจสอบชื่อซ้ำ
	var existName models.Product
	if err := database.DB.Where("name = ?", product.Name).First(&existName).Error; err == nil {
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
