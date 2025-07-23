package post

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	return c.JSON(posts)
}
// Giải thích:
// c *fiber.Ctx: context của request, chứa cả thông tin về request và response.
// c.JSON(posts): trả về dữ liệu JSON của slice posts.

func CreatePost(c *fiber.Ctx) error{
	var newPost Post
	if err := c.BodyParser(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Giải thích:
	// c.BodyParser(&newPost): phân tích body của request và gán vào biến newPost.
	// Nếu JSON sai định dang, trả về lỗi 400 Bad Request.
	newPost.ID = nextID
	nextID++
	posts = append(posts, newPost)

	return c.Status(fiber.StatusCreated).JSON(newPost)		// Trả về JSON của bài viết vừa tạo với status code 201 Created
}

func DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")	// Lấy ID từ path param (lấy :id trong URL, ví dụ: /posts/1)
	id, err := strconv.Atoi(idParam)	// Chuyển đổi ID từ string sang int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	for i, p := range posts {
		if p.ID == id {
			posts = append (posts[:i], posts[i+1:]...) // Kĩ thuật cắt slice để xóa phần tử i
			return c.SendStatus(fiber.StatusNoContent) 	// Trả về status code 204 No Content khi xóa thành công
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	})
}