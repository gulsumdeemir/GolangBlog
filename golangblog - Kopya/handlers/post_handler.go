package handlers

import (
	"fmt"
	"golangblog/database"
	"golangblog/models"
	"time"

	"github.com/gofiber/fiber/v2"
)


func GetPost(c *fiber.Ctx) error {
    postID := c.Query("id")
    categoryID := c.Query("categoryID")

    var posts []models.Post
    if postID != "" {
        var post models.Post
        result := database.DB.First(&post, postID)
        if result.Error != nil {
            return c.Status(404).SendString("Yazı bulunamadı")
        }
        posts = append(posts, post)
    } else if categoryID != "" {
        result := database.DB.Where("categoryID = ?", categoryID).Find(&posts)
        if result.Error != nil {
            return c.Status(500).SendString("Yazılar alınamadı: " + result.Error.Error())
        }
    } else {
        result := database.DB.Find(&posts)
        if result.Error != nil {
            return c.Status(500).SendString("Yazılar alınamadı: " + result.Error.Error())
        }
    }

    var postsWithTags []map[string]interface{}
    for _, post := range posts {
        var tags []models.Tag
        result := database.DB.Joins("JOIN PostTags ON PostTags.tagID = Tags.tagID").
            Where("PostTags.postID = ?", post.PostID).Find(&tags)
        if result.Error != nil {
            return c.Status(500).SendString("Etiketler alınamadı: " + result.Error.Error())
        }

        postMap := map[string]interface{}{
            "postID":          post.PostID,
            "postTitle":       post.PostTitle,
            "postDescription": post.PostDescription,
            "datee":           post.Datee,
            "userID":          post.UserID,
            "categoryID":      post.CategoryID,
            "tags":            tags,
        }

        postsWithTags = append(postsWithTags, postMap)
    }

    if postID != "" && len(postsWithTags) == 1 {
        return c.JSON(postsWithTags[0])
    }
    return c.JSON(postsWithTags)
}

func CreatePost(c *fiber.Ctx) error {
    var postInput struct {
        models.Post
        Tags []int `json:"tags"`
    }

    if err := c.BodyParser(&postInput); err != nil {
        fmt.Println("BodyParser hatası:", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz İstek"})
    }

    loggedInUser := c.Locals("userID")
    if loggedInUser == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    userID, ok := loggedInUser.(int)
    if !ok {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID'si"})
    }

    postInput.Post.UserID = userID
    postInput.Post.Datee = time.Now()

    if err := database.DB.Create(&postInput.Post).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Yazı oluşturulamadı"})
    }

    for _, tagID := range postInput.Tags {
        postTag := models.PostTag{PostID: postInput.Post.PostID, TagID: tagID}
        if err := database.DB.Create(&postTag).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostTag oluşturulamadı"})
        }
    }

    return c.JSON(fiber.Map{"message": "Yazı başarıyla oluşturuldu"})
}

func UpdatePost(c *fiber.Ctx) error {
    id := c.Params("id")
    var postInput struct {
        models.Post
        Tags []int `json:"tags"`
    }

    result := database.DB.First(&postInput.Post, id)
    if result.Error != nil {
        return c.Status(404).SendString("Yazı bulunamadı")
    }

    loggedInUser := c.Locals("userID")
    if loggedInUser == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    userID, ok := loggedInUser.(int)
    if !ok || postInput.Post.UserID != userID {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    if err := c.BodyParser(&postInput); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    database.DB.Save(&postInput.Post)

 
    database.DB.Where("postID = ?", postInput.Post.PostID).Delete(&models.PostTag{})

  
    for _, tagID := range postInput.Tags {
        postTag := models.PostTag{PostID: postInput.Post.PostID, TagID: tagID}
        if err := database.DB.Create(&postTag).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostTag oluşturulamadı"})
        }
    }

    return c.Status(200).JSON(postInput.Post)
}


func DeletePost(c *fiber.Ctx) error {
    id := c.Params("id")
    var post models.Post
    result := database.DB.First(&post, id)
    if result.Error != nil {
        return c.Status(404).SendString("Yazı bulunamadı")
    }

    loggedInUser := c.Locals("userID")
    if loggedInUser == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    userID, ok := loggedInUser.(int)
    if !ok || post.UserID != userID {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    database.DB.Delete(&post)
    return c.Status(200).SendString("Yazı silindi")
}


func GetUserPosts(c *fiber.Ctx) error {
    userID := c.Locals("userID")
    if userID == nil {
        return c.Status(400).SendString("Kullanıcı ID'si bulunamadı")
    }

    var posts []models.Post 
    result := database.DB.Where("userID = ?", userID).Find(&posts)
    if result.Error != nil {
        return c.Status(500).SendString("Yazılar alınamadı: " + result.Error.Error())
    }

    return c.JSON(posts)
}