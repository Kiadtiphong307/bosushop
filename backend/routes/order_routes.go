package routes

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router fiber.Router) {
	order := router.Group("/orders")

	order.Post("/", middleware.JWTAuth(""), controller.CreateOrder)
	order.Get("/", middleware.JWTAuth(""), controller.GetMyOrders)

	admin := router.Group("/admin/orders")
	admin.Get("/", middleware.JWTAuth("admin"), controller.GetAllOrders)
}
