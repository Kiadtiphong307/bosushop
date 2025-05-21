package controller

import (
	"backend/database"
	"backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// แสดงคูปองทั้งหมด
func GetAllCoupons(c *fiber.Ctx) error {
	var coupons []models.Coupon
	if err := database.DB.Order("created_at desc").Find(&coupons).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถโหลดคูปองได้"})
	}
	return c.JSON(coupons)
}

// สร้างคูปองใหม่
func CreateCoupon(c *fiber.Ctx) error {
	var input struct {
		Code            string    `json:"code"`
		DiscountPercent int       `json:"discount_percent"`
		MaxUsage        int       `json:"max_usage"`
		ExpireAt        time.Time `json:"expire_at"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	if input.Code == "" || input.DiscountPercent <= 0 || input.DiscountPercent > 100 || input.MaxUsage <= 0 || input.ExpireAt.Before(time.Now()) {
		return c.Status(400).JSON(fiber.Map{"error": "ข้อมูลคูปองไม่ถูกต้อง"})
	}

	coupon := models.Coupon{
		Code:            input.Code,
		DiscountPercent: input.DiscountPercent,
		MaxUsage:        input.MaxUsage,
		ExpireAt:        input.ExpireAt,
		UsedCount:       0,
	}

	if err := database.DB.Create(&coupon).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถสร้างคูปองได้"})
	}
	return c.Status(201).JSON(coupon)
}

// แก้ไขคูปอง
func UpdateCoupon(c *fiber.Ctx) error {
	id := c.Params("id")
	var coupon models.Coupon
	if err := database.DB.First(&coupon, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "ไม่พบคูปอง"})
	}

	var input struct {
		DiscountPercent int       `json:"discount_percent"`
		MaxUsage        int       `json:"max_usage"`
		ExpireAt        time.Time `json:"expire_at"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}

	coupon.DiscountPercent = input.DiscountPercent
	coupon.MaxUsage = input.MaxUsage
	coupon.ExpireAt = input.ExpireAt

	if err := database.DB.Save(&coupon).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ไม่สามารถอัปเดตคูปองได้"})
	}
	return c.JSON(coupon)
}

// ลบคูปอง
func DeleteCoupon(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Coupon{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ลบคูปองไม่สำเร็จ"})
	}
	return c.JSON(fiber.Map{"message": "ลบคูปองสำเร็จ"})
}
