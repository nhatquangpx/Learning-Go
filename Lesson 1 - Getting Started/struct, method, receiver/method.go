// Method là một hàm gắn với một kiểu dữ liệu, thường là struct
// Ví dụ: Thay vì viết hàm như sau:
func printPost(p Post){
	fmt.Println(p.Title)
}
// Chúng ta có thể viết hàm này như một phương thức của struct Post
func (p Post) Print() {
	fmt.Println(p.Title)
}

// Giải thích:
// - (p Post) là receiver, tức là đối tượng mà phương thức này sẽ làm việc với nó.
// - Print là tên của phương thức.