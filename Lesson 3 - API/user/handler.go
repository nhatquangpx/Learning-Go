package user

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("my_secret_key")

func Register(c *fiber.Ctx) error {
	var req User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	req.ID = nextID
	nextID++
	req.Password = string(hashed)
	Users = append(Users, req)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
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
	found := false
	for _, u:= range Users {
		if u.Username == req.Username {
			user = u
			found = true
			break
		}
	}

	if !found || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
	// user.Password đã được băm sẵn khi đăng ký, dùng bcrypt.GenerateFromPassword()
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
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