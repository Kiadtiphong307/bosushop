package routes

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
	products := router.Group("/products")

	products.Get("/", controller.GetAllProducts)
	products.Get("/:slug", controller.GetProductBySlug)

	// Protected admin routes
	products.Post("/", middleware.JWTAuth("admin"), controller.CreateProduct)
	products.Put("/:slug", middleware.JWTAuth("admin"), controller.UpdateProduct)
	products.Delete("/:slug", middleware.JWTAuth("admin"), controller.DeleteProduct)
}
