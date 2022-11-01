package main

import (
	"learning/golang-gormdatabase/database"
	"learning/golang-gormdatabase/routes"
	"learning/golang-gormdatabase/database/migration"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//init database
	database.InitDb()

	//migrasi database
	migration.RunMigration()

	//init fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat belajar",
			"city": "Depok",
		})
	})

	routes.RouteInit(app)
	app.Listen(":8000")
}