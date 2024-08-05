package handlers

import (
	"golangblog/database"
	"golangblog/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kullanıcı Bulunamadı")
	}


	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&user)

	return c.Status(200).JSON(user)
}


func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kullanıcı Bulunamadı")
	}
	database.DB.Delete(&user)

	return c.Status(200).SendString("Kullanıcı Silindi")
}

func UpdateProfile(c *fiber.Ctx) error {
	loggedInUser := c.Locals("userID")
	if loggedInUser == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Kullanıcı giriş yapmamış"})
	}

	var dbUser models.User
	result := database.DB.First(&dbUser, loggedInUser) 
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Kullanıcı Bulunamadı"})
	}

	var updateData models.User
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Geçersiz istek"})
	}

	dbUser.UserName = updateData.UserName
	dbUser.EMail = updateData.EMail
	if updateData.UserPassword != "" {
		dbUser.UserPassword = updateData.UserPassword
	}

	result = database.DB.Save(&dbUser) 
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Kullanıcı güncellenemedi"})
	}

	return c.Status(200).JSON(dbUser) 
}
func GetProfile(c *fiber.Ctx) error {
	loggedInUser := c.Locals("userID")
	if loggedInUser == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Kullanıcı giriş yapmamış"})
	}

	var user models.User
	if err := database.DB.First(&user, loggedInUser).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Kullanıcı Bulunamadı"})
	}

	return c.JSON(user)
}





  

