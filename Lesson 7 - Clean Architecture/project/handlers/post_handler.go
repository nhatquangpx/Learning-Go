// Xử lý HTTP request (controller)

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lesson-7/repositories"
)

func GetPost(c *fiber.Ctx) error {
	posts, err := repositories.GetAllPosts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.JSON(posts)
}