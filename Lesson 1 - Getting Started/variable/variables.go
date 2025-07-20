package main

import "fmt"

func main() {
	// Có 2 cách để khai báo biến
	// Cách 1: Khai báo biến với từ khóa var
	// var variablename type = value
	var name string = "Nguyen Van A"
	// Cach 2: Khai báo biến với từ khóa :=
	// variablename := value
	name2 := "Nguyen Van B"
	// Trong cách 2, Go sẽ tự động suy luận kiểu dữ liệu của biến, bắt buộc phải gán giá trị cho biến khi khai báo
	fmt.Println("Name 1:", name)
	fmt.Println("Name 2:", name2)
	// Nếu như khai báo biến với từ khóa var, ta có thể không gán giá trị cho biến
	// Khi đó, Go sẽ tự động gán giá trị mặc định cho biến
	// Giá trị mặc định của biến kiểu string là ""
	// Giá trị mặc định của biến kiểu int là 0
	// Giá trị mặc định của biến kiểu bool là false
	// Giá trị mặc định của biến kiểu float64 là 0.0
	var name3 string
	fmt.Println("Name 3:", name3)
	var age int
	fmt.Println("Age:", age)
	var isStudent bool
	fmt.Println("Is Student:", isStudent)
	var height float64
	fmt.Println("Height:", height)
	// Gán giá trị sau khi khai báo
	name3 = "Nguyen Van C"
	fmt.Println("Name 3 after assignment:", name3)
	// Khai báo nhiều biến cùng lúc
	var name4, name5 string = "Nguyen Van D", "Nguyen Van E"
	fmt.Println("Name 4:", name4)
	fmt.Println("Name 5:", name5)
	// Có thể khai báo kiểu dữ liệu khác nhau trong cùng một câu lệnh
	var name6, age2, isStudent2 = "Nguyen Van F", 20, true
	fmt.Println("Name 6:", name6)
	fmt.Println("Age 2:", age2)
	fmt.Println("Is Student 2:", isStudent2)
	// Khai báo biến trong một khối
	var (
		name7  string = "Nguyen Van G"
		age3   int    = 25
		isStudent3 bool = false
	)
	fmt.Println("Name 7:", name7)
	fmt.Println("Age 3:", age3)
	fmt.Println("Is Student 3:", isStudent3)

}

// Sự khác biết giữa cách khai báo biến với từ khóa var và := là:
// - Biến khai báo bằng var có thể sử dụng trong hoặc ngoài hàm, trong khi biến khai báo bằng := chỉ có thể sử dụng trong hàm.
// - Khai báo và gán giá trị bằng var có thể tách rời, trong khi khai báo và gán giá trị bằng := phải thực hiện cùng lúc.

// Quy tắc đặt tên biến:
// - Tên biến phải bắt đầu bằng chữ cái hoặc dấu gạch dưới (_).
// - Tên biến không được chứa khoảng trắng.
// - Tên biến không được chứa ký tự đặc biệt, ngoại trừ dấu gạch dưới (_).
// - Tên biến không được trùng với từ khóa của ngôn ngữ Go.
// - Tên biến có phân biệt chữ hoa và chữ thường, ví dụ: name và Name là hai biến khác nhau.
// - Tên biến không thể bắt đầu bằng số, ví dụ: 1name là tên biến không hợp lệ.
// - Tên biến nhiều từ:
//  - Có thể sử dụng dấu gạch dưới (_) để phân tách các từ, ví dụ: my_variable. (snake_case)
//  - Có thể viết liền không dấu gạch dưới, chữ cái đầu từ đầu tiên viết thường, chữ cái đầu của các từ sau viết hoa, ví dụ: myVariable. (camelCase)
//  - Có thể viết liền không dấu gạch dưới, chữ cái đầu của từ đầu tiên viết hoa, chữ cái đầu của các từ sau viết hoa, ví dụ: MyVariable. (PascalCase)