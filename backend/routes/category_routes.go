package routes


import (
	"backend/controller"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(router fiber.Router) {
	admin := router.Group("/admin/categories", middleware.JWTAuth("admin"))
	admin.Get("/", controller.GetAllCategories)
	admin.Post("/", controller.CreateCategory)
	admin.Delete("/:slug", controller.DeleteCategory)
}
