package controller

import (
	"blogs/database"
	"blogs/model"

	"github.com/gofiber/fiber/v2"
)

func GetBlogs(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "BlogGet",
	}
	db := database.DbConn
	var records model.Blogs
	db.Find(&records)
	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)
}
