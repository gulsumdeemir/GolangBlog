package main

import (
	"golangblog/database"
	"golangblog/handlers"
	"golangblog/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	
	app.Static("/", "./public")

	database.Connect()


	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", middleware.Protected(), handlers.UpdateUser)
	app.Delete("/users/:id", middleware.Protected(), handlers.DeleteUser)


	app.Get("/category", handlers.GetCategory)
	app.Post("/category", handlers.CreateCategory)
	app.Put("/category/:id", middleware.Protected(), handlers.UpdateCategory)
	app.Delete("/category/:id", middleware.Protected(), handlers.DeleteCategory)


	app.Get("/post/:id", handlers.GetPost) 
	app.Get("/post", handlers.GetPost)
	app.Post("/post", middleware.Protected(), handlers.CreatePost)
	app.Put("/post/:id", middleware.Protected(), handlers.UpdatePost)
	app.Delete("/post/:id", middleware.Protected(), handlers.DeletePost)




	app.Get("/comment", handlers.GetComments)       
	app.Post("/comment", middleware.Protected(), handlers.CreateComment) 
	app.Put("/comment/:id", middleware.Protected(), handlers.UpdateComment) 
	app.Delete("/comment/:id", middleware.Protected(), handlers.DeleteComment) 


	app.Get("/tag", handlers.GetTag)
	app.Post("/tag", handlers.CreateTag)
	app.Put("/tag/:id", middleware.Protected(), handlers.UpdateTag)
	app.Delete("/tag/:id", middleware.Protected(), handlers.DeleteTag)


	app.Get("/posttag", handlers.GetPostTag)
	app.Post("/posttag", middleware.Protected(), handlers.CreatePostTag)
	app.Put("/posttag/:id", middleware.Protected(), handlers.UpdatePostTag)
	app.Delete("/posttag/:id", middleware.Protected(), handlers.DeletePostTag)

	
	app.Get("/save", handlers.GetSave)
	app.Post("/save", middleware.Protected(), handlers.CreateSave)
	app.Put("/save/:id", middleware.Protected(), handlers.UpdateSave)
	app.Delete("/save/:id", middleware.Protected(), handlers.DeleteSave)

	app.Post("/login", handlers.LoginHandler)
	app.Post("/register", handlers.RegisterHandler)

	app.Post("/profile/update", middleware.Protected(), handlers.UpdateProfile)
	app.Get("/profile", middleware.Protected(), handlers.GetProfile)

	app.Get("/user/:userID/posts", middleware.Protected(), handlers.GetUserPosts)
	app.Get("/user/:userID/saves", middleware.Protected(), handlers.GetSavedPosts)

	
	log.Fatal(app.Listen(":8080"))
}
