package handlers

import (
	"golangblog/database"
	"golangblog/models"
	"log"

	"github.com/gofiber/fiber/v2"
)


func GetComments(c *fiber.Ctx) error {
	postID := c.Query("postID")
	var comments []models.Comment
	result := database.DB.Where("postID = ?", postID).Find(&comments)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(comments)
}


func CreateComment(c *fiber.Ctx) error {
    comment := new(models.Comment)
    userID := c.Locals("userID").(int)

    if err := c.BodyParser(comment); err != nil {
        log.Println("Body parsing error:", err)
        return c.Status(400).SendString(err.Error())
    }

    comment.UserID = userID 
    log.Printf("Comment: %+v\n", comment)

    if comment.PostID == nil {
        return c.Status(400).SendString("postID is required")
    }

    result := database.DB.Create(comment)
    if result.Error != nil {
        log.Println("Database error:", result.Error)
        return c.Status(500).SendString(result.Error.Error())
    }
    return c.Status(201).JSON(comment)
}
func UpdateComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	result := database.DB.First(&comment, id)
	if result.Error != nil {
		return c.Status(404).SendString("Yorum bulunamadı")
	}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&comment)
	return c.Status(200).JSON(comment)
}


func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	result := database.DB.First(&comment, id)
	if result.Error != nil {
		return c.Status(404).SendString("Yorum Bulunamadı")
	}
	database.DB.Delete(&comment)
	return c.Status(200).SendString("Yorum Silindi")
}


