package post

import "fmt"

func ShowMenu() {
	fmt.Println("\n Quản lý bài viết")
	fmt.Println("1. Thêm bài viết")
	fmt.Println("2. Hiển thị tất cả bài viết")
	fmt.Println("3. Sửa bài viết")
	fmt.Println("4. Xóa bài viết")
	fmt.Println("5. Thoát")
	fmt.Print("Chọn một tùy chọn: ")

	var choice int
	fmt.Scanln(&choice)
	
	switch choice {
	case 1:
		addPost()
	case 2:
		displayPosts()
	case 3:
		editPost()
	case 4:
		deletePost()
	case 5:
		fmt.Println("Tạm bịc nho.")
		return
	default:
		fmt.Println("Lựa chọn không hợp lệ.")
	}
}

func addPost() {
	var title, content string
	fmt.Print("Nhập tiêu đề bài viết: ")
	fmt.Scanln(&title)
	fmt.Print("Nhập nội dung bài viết: ")
	fmt.Scanln(&content)

	post := Post{id: nextId, title: title, content: content}
	Posts = append(Posts, post)
	nextId++
	fmt.Println("Bài viết đã được thêm thành công.")
}

func displayPosts() {
	if len(Posts) == 0 {
		fmt.Println("Không có bài viết nào để hiển thị.")
		return
	}
	for _, p := range Posts {
		fmt.Printf("ID: %d\nTiêu đề: %s\nNội dung: %s\n--- \n", p.id, p.title, p.content)
	}
}

func editPost() {
	if len(Posts) == 0 {
		fmt.Println("Không có bài viết nào để sửa.")
		return
	}

	var id int
	fmt.Print("Nhập ID bài viết cần sửa: ")
	fmt.Scanln(&id)

	index := -1
	for i, p := range Posts {
		if p.id == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Bài viết không tồn tại.")
		return
	}

	var newTitle, newContent string
	fmt.Print("Nhập tiêu đề mới: ")
	fmt.Scanln(&newTitle)
	fmt.Print("Nhập nội dung mới: ")
	fmt.Scanln(&newContent)

	Posts[index].title = newTitle
	Posts[index].content = newContent
	fmt.Println("Bài viết đã được cập nhật thành công.")
}

func deletePost() {
	if len(Posts) == 0 {
		fmt.Println("Không có bài viết nào để xóa.")
		return
	}

	var id int
	fmt.Print("Nhập ID bài viết cần xóa: ")
	fmt.Scanln(&id)

	index := -1
	for i, p := range Posts {
		if p.id == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Bài viết không tồn tại.")
		return
	}

	Posts = append(Posts[:index], Posts[index+1:]...)
	fmt.Println("Bài viết đã được xóa thành công.")
}