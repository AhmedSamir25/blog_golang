package router

import (
	"blogs/controller"

	"github.com/gofiber/fiber/v2"
)

func StepRouters(app *fiber.App) {
	app.Get("/", controller.GetBlogs)
	app.Post("/", controller.CreateBlog)
}
