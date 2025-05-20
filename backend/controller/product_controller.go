package controller

import (
	"backend/database"
	"backend/models"
	"backend/validation"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

// ดึงข้อมูลสินค้าทั้งหมด
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := database.DB.Order("created_at desc").Find(&products).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถดึงข้อมูลสินค้าได้"})
	}
	return c.JSON(products)
}

// ดึงข้อมูลสินค้าตาม ID
func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "ไม่พบสินค้า"})
	}
	return c.JSON(product)
}

// สร้างสินค้าใหม่
func CreateProduct(c *fiber.Ctx) error {
	var input validation.ProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid format"})
	}

	validationErrors := validation.ValidateProductInput(input)
	if validationErrors != nil {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
		Available:   input.Available,
		Stock:       input.Stock,
	}

	if err := services.CreateProduct(&product); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(product)
}

// แก้ไขข้อมูลสินค้า
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "ไม่พบสินค้า"})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถอัปเดตสินค้าได้"})
	}
	return c.JSON(product)
}

// ลบสินค้า
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ลบสินค้าไม่สำเร็จ"})
	}
	return c.JSON(fiber.Map{"message": "ลบสินค้าสำเร็จ"})
}
