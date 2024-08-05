package handlers

import (
	"golangblog/database"
	"golangblog/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "your_secret_key"

func LoginHandler(c *fiber.Ctx) error {
    var login models.User

    if err := c.BodyParser(&login); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz İstek"})
    }

    var user models.User
    if err := database.DB.Where("eMail = ? AND userName = ?", login.EMail, login.UserName).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz Bilgi"})
    }

    if user.UserPassword != login.UserPassword {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz Bilgi"})
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": user.UserID,
        "exp":    time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte(SecretKey))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token oluşturulamadı."})
    }

    return c.JSON(fiber.Map{
        "message":   "Giriş Başarılı",
        "token":     tokenString,
        "userName":  user.UserName,
        "eMail":     user.EMail,
    })
}

