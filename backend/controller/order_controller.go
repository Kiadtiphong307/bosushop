package controller

import (
	"backend/services"
	"backend/validation"
	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	var input validation.OrderInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}

	if errors := validation.ValidateOrderInput(input); errors != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errors})
	}

	userID := c.Locals("userID").(uint)
	var coupon *string
	if input.CouponCode != "" {
		coupon = &input.CouponCode
	}

	order, err := services.CreateOrder(userID, input.ProductID, coupon)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(order)
}

func GetMyOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	orders, err := services.GetOrdersByUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "โหลดคำสั่งซื้อไม่สำเร็จ"})
	}
	return c.JSON(orders)
}

func GetAllOrders(c *fiber.Ctx) error {
	orders, err := services.GetAllOrders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "โหลดคำสั่งซื้อทั้งหมดไม่สำเร็จ"})
	}
	return c.JSON(orders)
}
