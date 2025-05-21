package services

import (
	"backend/database"
	"backend/models"
	"errors"

	"github.com/gosimple/slug"
)

// ดึงสินค้าทั้งหมด
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := database.DB.
		Preload("Category").
		Order("created_at desc").
		Find(&products).Error
	return products, err
}

// ดึงสินค้าตาม slug (แยกจาก fiber)
func GetProductBySlug(slug string) (*models.Product, error) {
	var product models.Product
	if err := database.DB.
		Preload("Category").
		Where("slug = ?", slug).
		First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// สร้างสินค้าใหม่
func CreateProduct(product *models.Product) error {
	product.Slug = slug.MakeLang(product.Name, "th")

	var existSlug models.Product
	if err := database.DB.Where("slug = ?", product.Slug).First(&existSlug).Error; err == nil {
		return errors.New("มีสินค้าที่ใช้ slug เดียวกันแล้ว")
	}

	var existName models.Product
	if err := database.DB.Where("name = ?", product.Name).First(&existName).Error; err == nil {
		return errors.New("ชื่อสินค้านี้มีอยู่แล้ว")
	}

	return database.DB.Create(product).Error
}

// แก้ไขสินค้า
func UpdateProduct(id uint, data *models.Product) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil, err
	}

	product.Name = data.Name
	product.Description = data.Description
	product.Price = data.Price
	product.ImageURL = data.ImageURL
	product.Available = data.Available
	product.Stock = data.Stock
	product.CategoryID = data.CategoryID
	product.Slug = slug.MakeLang(data.Name, "th")

	if err := database.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// ลบสินค้า
func DeleteProduct(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}
