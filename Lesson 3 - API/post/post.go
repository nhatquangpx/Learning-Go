package post

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts = []Post{}
var nextID = 1				// ID tự động tăng cho bài viết mới

// Dữ liệu sẽ chỉ lưu tạm trong RAM, khi restart file sẽ mất
