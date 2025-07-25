package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	DB = db
	log.Println("PostgreSQL database connected successfully!")
}

// Giải thích:
// - Hàm ConnectDatabase() kết nối đến PostgreSQL sử dụng GORM.
// - Sprintf để định dạng chuỗi DSN từ biến môi trường.
// - gorm.Open() mở kết nối đến cơ sở dữ liệu.
// - postgres.Open(dsn) sử dụng driver PostgreSQL với chuỗi kết nối DSN.
// - &gorm.Config{} là cấu hình mặc định của GORM.
// - Nếu kết nối thất bại, log.Fatal sẽ in ra lỗi và dừng chương trình.
// - Nếu thành công, in ra thông báo kết nối thành công.
// - Biến DB sẽ lưu trữ kết nối để sử dụng trong các phần khác của ứng dụng.