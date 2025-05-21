package main

import (
	"github.com/gofiber/fiber/v2"

	"backend/database"
	"backend/routes"
	"backend/seed"
	"fmt"
)

func main() {
	app := fiber.New()

	// เชื่อมต่อฐานข้อมูล
	database.InitDatabase()

	// เรียกใช้ routes ทั้งหมด
	routes.AuthRoutes(app)     // สำหรับการสมัครสมาชิกและเข้าสู่ระบบ
	routes.OrderRoutes(app)    // สำหรับการสร้างคำสั่งซื้อ
	routes.ProductRoutes(app)  // สำหรับการสร้างสินค้า
	routes.CategoryRoutes(app) // สำหรับการสร้างหมวดหมู่สินค้า Admin
	routes.CouponRoutes(app)   // สำหรับการสร้างคูปอง

	// เรียกใช้ seed ข้อมูล
	seed.InitSeeder(database.DB)


	app.Listen(":8080")
}
