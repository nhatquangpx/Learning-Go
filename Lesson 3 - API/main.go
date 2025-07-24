// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type Post struct {
// 	ID int	`json:"id"`
// 	Title string `json:"title"`
// 	Content string `json:"content"`
// }

// var posts = []Post{
// 	{ID: 1, Title: "First Post", Content: "This is the content of the first post."},
// 	{ID:2, Title: "Second Post", Content: "This is the content of the second post."},
// }

// func getPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(posts)
// }

// func main() {
// 	http.HandleFunc("/posts", getPosts)
// 	fmt.Println("Server is running on port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// Giải thích:
// 1. Định nghĩa một struct `Post` để mô tả cấu trúc của bài viết.
// 2. Tạo một slice `posts` chứa các bài viết mẫu.
// 3. Hàm `getPosts` xử lý yêu cầu GET đến endpoint `/posts`, trả về danh sách bài viết dưới dạng JSON.
// 4. Trong hàm `main`, thiết lập một HTTP server lắng nghe trên cổng 8080 và đăng ký hàm xử lý cho endpoint `/posts`.
// 5. Sử dụng `json.NewEncoder` để mã hóa slice `posts` thành JSON và gửi về cho client.
// 6. Sử dụng `log.Fatal` để ghi log nếu có lỗi xảy ra
// Trong hàm getPosts, tham số `w` là đối tượng `http.ResponseWriter` dùng để gửi phản hồi về client
// `r` là đối tượng `*http.Request` chứa thông tin về yêu cầu HTTP.

// Giới hạn của net/http:
// Viết router thủ công, không hỗ trợ path param (ví dụ: /posts/:id)
// Không dễ cấu trúc nhiều file
// Không hỗ trợ middleware, logging, body parser, validator

// => Dùng Fiber để giải quyết các vấn đề này
package main

import (
	"go-http-api/post"
	"go-http-api/user"
	"go-http-api/middleware"
	"github.com/gofiber/fiber/v2"		// Fiber là framework HTTP hiện đại và nhẹ, tương tự Express của Node.js
)

func main() {
	app := fiber.New()		 		 	// Trả về một *App, đại diện cho HTTP server

	app.Post("/register", user.Register)
	app.Post("/login", user.Login)

	api := app.Group("/posts", middleware.RequireAuth)

	app.Get("/", post.GetPosts)	// GET tất cả bài viết
	app.Post("/", post.CreatePost)	// POST bài viết mới
	app.Delete("/:id", post.DeletePost) // DELETE bài viết theo ID
	app.Get("/:id", post.GetPostByID) // GET bài viết theo ID
	app.Put("/:id", post.UpdatePost) // PUT cập nhật bài viết theo ID
	app.Patch(":id", post.PatchPostTitle) // PATCH cập nhật title bài viết theo ID
	app.Listen(":3000")				// Bắt đầu chạy server ở cổng 3000
}
