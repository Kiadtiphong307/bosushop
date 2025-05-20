package main

import (


	"github.com/gofiber/fiber/v2"

	"backend/database"
	"backend/routes"
)

func main() {
	app := fiber.New()

	// เชื่อมต่อฐานข้อมูล
	database.InitDatabase()

	// เรียกใช้ routes ทั้งหมด
	routes.AuthRoutes(app) // routes สำหรับการสมัครสมาชิกและเข้าสู่ระบบ

	app.Listen(":8080")
}
