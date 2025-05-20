package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"backend/database"
)

func main() {
	app := fiber.New()

	// เชื่อมต่อฐานข้อมูล
	database.InitDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}
