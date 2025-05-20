
package controller

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"backend/database"
)

// สมัครสมาชิก 
func Register(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	input.Role = "user" // ป้องกัน user สมัครเป็น admin

	if err := services.RegisterUser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Register failed"})
	}

	return c.JSON(fiber.Map{
		"message": "Register success",
		"user": input,
	})
}

// เข้าสู่ระบบ
func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := services.Authenticate(input.Email, input.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)
	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

// แสดงข้อมูลสมาชิก
func Profile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}
