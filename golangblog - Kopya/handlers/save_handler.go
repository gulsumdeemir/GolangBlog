package handlers

import (
	"golangblog/database"
	"golangblog/models"

	"github.com/gofiber/fiber/v2"
)

func GetSave(c *fiber.Ctx) error {
	var save []models.Save
	result := database.DB.Find(&save)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(save)
}


func CreateSave(c *fiber.Ctx) error {
	save := new(models.Save)
	if err := c.BodyParser(save); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	result := database.DB.Create(&save)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(201).JSON(save)
}


func UpdateSave(c *fiber.Ctx) error {
	id := c.Params("id")
	var save models.Save
	result := database.DB.First(&save, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kaydedilme bulunamadı")
	}
	if err := c.BodyParser(&save); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&save)
	return c.Status(200).JSON(save)
}


func DeleteSave(c *fiber.Ctx) error {
	id := c.Params("id")
	var save models.Save
	result := database.DB.First(&save, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kaydedilme Bulunamadı")
	}
	database.DB.Delete(&save)
	return c.Status(200).SendString("Kaydedilme Silindi")
}

func GetSavedPosts(c *fiber.Ctx) error {
    userID := c.Locals("userID")
    if userID == nil {
        return c.Status(400).SendString("Kullanıcı ID'si bulunamadı")
    }

    userIDInt, ok := userID.(int)
    if !ok {
        return c.Status(400).SendString("Geçersiz kullanıcı ID'si")
    }

    var saves []models.Save
    result := database.DB.Where("userID = ?", userIDInt).Find(&saves)
    if result.Error != nil {
        return c.Status(500).SendString("Kaydedilen yazılar alınamadı: " + result.Error.Error())
    }

    var posts []models.Post
    for _, save := range saves {
        var post models.Post
        if err := database.DB.First(&post, save.PostID).Error; err != nil {
            return c.Status(500).SendString("Post detayları alınamadı: " + err.Error())
        }
        posts = append(posts, post)
    }

    return c.JSON(posts)
}







