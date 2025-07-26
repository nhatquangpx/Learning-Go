// Giao tiếp với cơ sở dữ liệu (CRUD)

package repositories

import (
	"lesson-7/models"
	"lesson-7/config"
)

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := config.DB.Find(&posts).error
	return posts, err
}