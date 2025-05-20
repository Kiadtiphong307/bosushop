package controller

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := database.DB.Order("created_at desc").Find(&products).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถดึงข้อมูลสินค้าได้"})
	}
	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "ไม่พบสินค้า"})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var input models.Product
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถสร้างสินค้าได้"})
	}
	return c.Status(201).JSON(input)
}

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

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ลบสินค้าไม่สำเร็จ"})
	}
	return c.JSON(fiber.Map{"message": "ลบสินค้าสำเร็จ"})
}
