package handlers

import (
	"golangblog/database"
	"golangblog/models"

	"github.com/gofiber/fiber/v2"
)

func GetPostTag(c *fiber.Ctx) error {
	var posttag []models.PostTag
	result := database.DB.Find(&posttag)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(posttag)
}

func CreatePostTag(c *fiber.Ctx) error {
	posttag := new(models.PostTag)
	if err := c.BodyParser(posttag); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	result := database.DB.Create(&posttag)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(201).JSON(posttag)
}

func UpdatePostTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var posttag models.PostTag
	result := database.DB.First(&posttag, id)
	if result.Error != nil {
		return c.Status(404).SendString("bulunamadı")
	}
	if err := c.BodyParser(&posttag); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&posttag)
	return c.Status(200).JSON(posttag)
}


func DeletePostTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var posttag models.PostTag
	result := database.DB.First(&posttag, id)
	if result.Error != nil {
		return c.Status(404).SendString("Bulunamadı")
	}
	database.DB.Delete(&posttag)
	return c.Status(200).SendString("Silindi")
}