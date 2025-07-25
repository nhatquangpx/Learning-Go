package post

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10" 
)

var validate = validator.New()
func GetPosts(c *fiber.Ctx) error {
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
	newPost.ID = nextID
	nextID++
	newPost.AuthorID = c.Locals("user_id").(int) 
	posts = append(posts, newPost)

	return c.Status(fiber.StatusCreated).JSON(newPost)		
}

func DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")	
	id, err := strconv.Atoi(idParam)	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	for i, p := range posts {
		if p.ID == id {
			posts = append (posts[:i], posts[i+1:]...) 
			return c.SendStatus(fiber.StatusNoContent) 	
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
	for i, p := range posts {
		if p.ID == id {
			updated.ID = id	
			updated.AuthorID = p.AuthorID 
			posts[i] = updated
			return c.JSON(updated)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	})
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

	for i, p := range posts {
		if p.ID == id {
			posts[i].Title = req.Title
			return c.JSON(posts[i])
		}
	 }

	 return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	 })
}