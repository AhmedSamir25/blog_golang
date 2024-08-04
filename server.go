package main

import (
	"blogs/database"
	"blogs/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnDB()
	app := fiber.New()

	router.StepRouters(app)
	app.Listen(":3001")
}
