
package controller

import (
	"backend/models"
	"backend/validation"
	"github.com/gofiber/fiber/v2"
	"backend/database"
	"backend/services"
	"backend/utils"
)

// สมัครสมาชิก
func Register(c *fiber.Ctx) error {
	var input validation.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid format"})
	}
	if err := validation.ValidateRegisterInput(input); err != nil {
		return c.Status(422).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := services.RegisterUser(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)
	return c.JSON(fiber.Map{"token": token, "user": user})
}

// เข้าสู่ระบบ
func Login(c *fiber.Ctx) error {
	var input validation.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validation.ValidateLoginInput(input); err != nil {
		return c.Status(422).JSON(fiber.Map{"error": err.Error()})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	token, _ := utils.GenerateJWT(user.ID, user.Role)

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
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
