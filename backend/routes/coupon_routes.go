package routes

import (
	"backend/controller"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func CouponRoutes(router fiber.Router) {
	admin := router.Group("/admin/coupons", middleware.JWTAuth("admin"))
	admin.Get("/", controller.GetAllCoupons)
	admin.Post("/", controller.CreateCoupon)
	admin.Put("/:id", controller.UpdateCoupon)
	admin.Delete("/:id", controller.DeleteCoupon)
}
