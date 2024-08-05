package controller

import (
	"blogs/database"
	"blogs/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetBlogs(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "BlogGet",
	}
	db := database.DbConn
	var records []model.Blogs
	db.Find(&records)
	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)
}
func CreateBlog(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Create Blog",
	}

	record := new(model.Blogs)
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request:", err)
		context["statusText"] = "bad"
		context["msg"] = "error in parsing"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}

	result := database.DbConn.Create(record)
	if result.Error != nil {
		log.Println("Error in saving data:", result.Error)
		context["statusText"] = "bad"
		context["msg"] = "error in save"
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(context)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(context)
}
