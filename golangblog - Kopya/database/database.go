package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:demir2001@tcp(127.0.0.1:3306)/blogsite?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veri tabanına bağlantı hatası:", err)
	}

	fmt.Println("Veri tabanına başarıyla bağlanıldı")
}
