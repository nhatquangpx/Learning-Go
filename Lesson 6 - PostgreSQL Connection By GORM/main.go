package main

import (
	"lesson-4/config"
	"lesson-4/routes"
	"github.com/gofiber/fiber/v2"
	_"lesson-4/docs"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)


// @title Fiber Blog API
// @version 1.0
// @description REST API with JWT Auth, written in Go using Fiber
// @host localhost:3000
// @BasePath /
func main() {
	config.LoadEnv() // Nạp biến môi trường từ .env
	config.ConnectDatabase() // Kết nối đến PostgreSQL
	app := fiber.New()

	// Đăng ký route Swagger trước khi chạy server
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	routes.Setup(app) // Thiết lập các route từ package routes

	port := config.GetEnv("PORT", "3000") // Lấy cổng từ biến môi trường, mặc định là 3000
	app.Listen(":" + port)

}