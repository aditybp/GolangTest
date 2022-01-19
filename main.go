package main

import (
	"golangtest/database"
	"golangtest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	database.ConnectionDB()
	routes.Setup(app)
	app.Listen(":8000")
}
