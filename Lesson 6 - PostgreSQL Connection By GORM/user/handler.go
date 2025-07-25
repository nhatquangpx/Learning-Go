package user

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"lesson-4/config"
	"time"
)

var jwtSecret = []byte("my_secret_key")

// Register godoc
// @Summary Đăng ký người dùng mới
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body user.User true "Thông tin người dùng"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var req User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	req.Password = string(hashed)
	if err := config.DB.Create(&req).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to save user",})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": req.ID,
	})
}

func Login(c *fiber.Ctx) error {
	var req User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	var user User
	if err := config.DB.Where("username = ?",req.Username).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	// JWT này sẽ được gắn vào các request sau này để xác thực người dùng

	signedToken, err := token.SignedString(jwtSecret)		// Ký token với secret key
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": signedToken,
	})
	// Gửi token cho client dưới dạng:
	//	{ "token": "<chuỗi JWT>" }
}