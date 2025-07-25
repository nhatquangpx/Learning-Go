package post

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10" 
	"lesson-4/config"
)

var validate = validator.New()
func GetPosts(c *fiber.Ctx) error {
	var posts []Post
	if err := config.DB.Find(&posts).Error; err !=nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve posts",
		})
	}
	return c.JSON(posts)
}

func CreatePost(c *fiber.Ctx) error{
	var newPost Post
	if err := c.BodyParser(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(newPost); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(), 
		})
	}
	newPost.AuthorID = c.Locals("user_id").(int) 
	if err := config.DB.Create(&newPost).Error; err != nil{
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newPost)		
}

func DeletePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var post Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	}
	if err := config.DB.Delete(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}
	return c.SendStatus(204)
}

func GetPostByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var post Post
	if err := config.DB.First(&post, id).Error; err!= nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Post not found",
		})
	}
	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var updated Post
	if err := c.BodyParser(&updated); err != nil {		
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := validate.Struct(updated); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var post Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	}
	post.Title = updated.Title
	post.Content = updated.Content
	if err := config.DB.Save(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}
	return c.JSON(post)
}

func PatchPostTitle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":" Invalid ID format",
		})
	}
	var req updateTitleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var post Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	}
	post.Title = req.Title
	if err := config.DB.Save(&post).Error; err != nil{
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update post title",
		})
	}
	 return c.JSON(post)
}