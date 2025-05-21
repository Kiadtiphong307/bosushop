package controller

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

// ดึงหมวดหมู่ทั้งหมด
func GetAllCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := database.DB.Order("created_at desc").Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถดึงหมวดหมู่ได้"})
	}
	return c.JSON(categories)
}

// เพิ่มหมวดหมู่ใหม่
func CreateCategory(c *fiber.Ctx) error {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&input); err != nil || input.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ชื่อหมวดหมู่จำเป็นต้องกรอก"})
	}

	category := models.Category{Name: input.Name}
	if err := database.DB.Create(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถเพิ่มหมวดหมู่ได้"})
	}
	return c.Status(201).JSON(category)
}

// ลบหมวดหมู่
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถลบหมวดหมู่ได้"})
	}
	return c.JSON(fiber.Map{"message": "ลบหมวดหมู่สำเร็จ"})
}
