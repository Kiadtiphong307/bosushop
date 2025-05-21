package routes

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
	products := router.Group("/products")

	products.Get("/", controller.GetAllProducts)
	products.Get("/:id", controller.GetProductByID)

	// Protected admin routes
	products.Post("/", middleware.JWTAuth("admin"), controller.CreateProduct)
	products.Put("/:id", middleware.JWTAuth("admin"), controller.UpdateProduct)
	products.Delete("/:id", middleware.JWTAuth("admin"), controller.DeleteProduct)
}
