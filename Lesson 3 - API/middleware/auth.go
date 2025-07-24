package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte("my_secret_key")

func RequireAuth(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid token",
		})
	}
	// c.Get("Authorization") lấy giá trị từ header: Authorization: <token>

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	// jwt.Parse sẽ giải mã token và trả về một jwt.Token
	// Hàm fun(t(t *jwt.Token) (interface{}, error) sẽ được gọi để lấy khóa bí mật dùng để xác thực token
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	// token.Claims chứa dữ liệu bên trong JWT, như:
	// {""user_id": 1, "exp": 1700000000}
	// Vì giá trị trong MapClaims là interface{} -> cần ép kiểu về float64 rồi chuyển sang int(do JSON luôn dùng float64 cho số)
	c.Locals("user_id", userID)
	// Bước này rất quan trọng:
	// Fiber cho phép gắn biến vào context request để sử dụng trong các handler sau này
	// Trong các handler, ta có thể lấy lại bằng: authorID := c.Locals("user_id").(int)
	return c.Next()			// Nếu token hợp lệ, tiếp tục xử lý request 
} 