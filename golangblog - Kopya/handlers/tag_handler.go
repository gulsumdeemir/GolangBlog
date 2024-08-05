package handlers

import (
	"golangblog/database"
	"golangblog/models"

	"github.com/gofiber/fiber/v2"
)

func GetTag(c *fiber.Ctx) error {
	var tag []models.Tag
	result := database.DB.Find(&tag)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(tag)
}
func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)
	if err := c.BodyParser(tag); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	result := database.DB.Create(&tag)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(201).JSON(tag)
}

func UpdateTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var tag models.Tag


	result := database.DB.First(&tag, id)
	if result.Error != nil {
		return c.Status(404).SendString("Etiket bulunamadı")
	}


	if err := c.BodyParser(&tag); err != nil {
		return c.Status(400).SendString(err.Error())
	}


	database.DB.Save(&tag)

	return c.Status(200).JSON(tag)
}


func DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var tag models.Tag

	result := database.DB.First(&tag, id)
	if result.Error != nil {
		return c.Status(404).SendString("Etiket Bulunamadı")
	}
	database.DB.Delete(&tag)

	return c.Status(200).SendString("Etiket Silindi")
}