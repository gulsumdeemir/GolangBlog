package handlers

import (
	"golangblog/database"
	"golangblog/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	user := new(models.User)


	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Kullanıcı verileri okunamadı")
	}

	user.Datee = new(time.Time)
	*user.Datee = time.Now()

	
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).SendString("Kullanıcı kaydedilirken bir hata oluştu: " + result.Error.Error())
	}

	return c.Status(201).JSON(user)
}
