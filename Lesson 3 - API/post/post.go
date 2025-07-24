package post

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title" validate:"required, min=3"`	// Thêm tag validate tối thiểu 3 ký tự
	Slug    string `json:"slug" validate:"required"`	// Thêm tag validate bắt buộc không được rỗng
	Content string `json:"content" validate:"required"`
	AuthorID int   `json:"author_id"` // Chuẩn bị cho phần login
}

var posts = []Post{}
var nextID = 1				// ID tự động tăng cho bài viết mới

// Dữ liệu sẽ chỉ lưu tạm trong RAM, khi restart file sẽ mất
