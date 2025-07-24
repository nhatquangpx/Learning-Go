package main

import (
	"lesson-4/config"
	"lesson-4/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv() // Nạp biến môi trường từ .env

	app := fiber.New()
	routes.Setup(app) // Thiết lập các route từ package routes

	port := config.GetEnv("PORT", "3000") // Lấy cổng từ biến môi trường, mặc định là 3000
	app.Listen(":" + port)
}