package post

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10" // Thư viện validate để kiểm tra dữ liệu
	// Thư viện phổ biến để kiểm tra dữ liệu client gửi lên có hợp lệ hay không
)

var validate = validator.New()
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
	if err := validate.Struct(newPost); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(), // Trả về lỗi validate nếu có
		})
	}
	newPost.ID = nextID
	nextID++
	newPost.AuthorID = 1 // Giả sử AuthorID là 1, sau này sẽ thay bằng ID của người dùng đã đăng nhập
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

func GetPostByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	for _, p := range posts {
		if p.ID == id {
			return c.JSON(p)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var updated Post
	if err := c.BodyParser(&updated); err != nil {		// Dùng BodyParser() để lấy JSON từ request body
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := validate.Struct(updated); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	for i, p := range posts {
		if p.ID == id {
			updated.ID = id	// Giữ nguyên ID cũ
			updated.AuthorID = p.AuthorID // Giữ nguyên AuthorID cũ
			posts[i] = updated
			return c.JSON(updated)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	})
}