package middleware

import (
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func JWTAuth(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if len(auth) < 7 || auth[:7] != "Bearer " {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}

		token := auth[7:]
		claims, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		role := claims["role"].(string)
		if requiredRole != "" && role != requiredRole {
			return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
		}

		c.Locals("userID", uint(claims["user_id"].(float64)))
		c.Locals("role", role)
		return c.Next()
	}
}
