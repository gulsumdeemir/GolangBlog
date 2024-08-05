package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Yetkilendirme gerekiyor")
		}

	
		token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

		claims, err := validateJWT(token)
		if err != nil {
			fmt.Println("JWT Doğrulama Hatası:", err)
			return c.Status(fiber.StatusUnauthorized).SendString("Geçersiz veya süresi dolmuş token")
		}

		userID, ok := claims["userID"].(float64)
		if !ok {
			fmt.Println("JWT Claims Hatası: userID bulunamadı")
			return c.Status(fiber.StatusUnauthorized).SendString("Geçersiz token claims")
		}

		fmt.Println("JWT doğrulandı, userID:", userID)
		c.Locals("userID", int(userID))

		return c.Next()
	}
}


func validateJWT(token string) (jwt.MapClaims, error) {
	secret := "your_secret_key" 
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("beklenmeyen imza yöntemi")
		}
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("JWT parse hatası:", err)
		return nil, err
	}

	if !parsedToken.Valid {
		fmt.Println("Geçersiz token")
		return nil, errors.New("geçersiz token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Claims hatası")
		return nil, errors.New("geçersiz token claims")
	}
	return claims, nil
}
