// Receiver là biến đứng trước tên hàm, đại diện cho đối tượng đang gọi hàm
// Cú pháp:
// func (receiver type) methodName(parameters) returnType {...}
// Ví dụ:
func (p Post) Summary() string {
	return p.Title + " - " + p.Content
}
// Gọi method:
p := Post{
	ID: 1,
	Title: "Hello World",
	Content: "This is my first post."
}
fmt.Println(p.Summary()) // In ra: Hello World - This is my first post.