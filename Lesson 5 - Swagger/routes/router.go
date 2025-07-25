package routes

import (
	"lesson-4/middleware"
	"lesson-4/post"
	"lesson-4/user"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", user.Register) // Route đăng ký người dùng
	app.Post("/login", user.Login)       // Route đăng nhập người dùng

	api := app.Group("/posts", middleware.RequireAuth)
	api.Get("/", post.GetPosts)	// GET tất cả bài viết
	api.Post("/", post.CreatePost)	// POST bài viết mới
	api.Delete("/:id", post.DeletePost) // DELETE bài viết theo ID
	api.Get("/:id", post.GetPostByID) // GET bài viết theo ID
	api.Put("/:id", post.UpdatePost) // PUT cập nhật bài viết theo ID
	api.Patch(":id", post.PatchPostTitle) // PATCH cập nhật title bài viết theo ID
}