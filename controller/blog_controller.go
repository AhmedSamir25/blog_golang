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
func UpdateBlog(c *fiber.Ctx) error {
	context := fiber.Map{"msg": "okupdate"}
	id := c.Params("id")
	var record model.Blogs
	database.DbConn.First(&record, id)
	if record.Id == 0 {
		log.Println("Record Not Found")
		return c.JSON(context)
	}
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in body reqoust")
	}
	reslut := database.DbConn.Save(record)
	if reslut.Error != nil {
		log.Println("Error in saving")
	}
	context["data"] = record
	c.Status(200)
	return c.JSON(context)
}

func DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var record model.Blogs

	result := database.DbConn.First(&record, id)
	if result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"msg":    "record not found",
			"status": "error",
		})
	}
	result = database.DbConn.Delete(&record)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"msg":    "error deleting record",
			"status": "error",
		})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"msg":    "delete is good",
		"status": "ok",
	})
}
