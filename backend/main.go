package main

import (


	"github.com/gofiber/fiber/v2"

	"backend/database"
	"backend/routes"
	"backend/seed"
)

func main() {
	app := fiber.New()

	// เชื่อมต่อฐานข้อมูล
	database.InitDatabase()

	// เรียกใช้ routes ทั้งหมด
	routes.Routes(app) // เรียกใช้ routes ทั้งหมด
	routes.OrderRoutes(app)
	routes.ProductRoutes(app)

	// เรียกใช้ seed ข้อมูล
	seed.InitSeeder(database.DB)

	app.Listen(":8080")
}
