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

}

// Sự khác biết giữa cách khai báo biến với từ khóa var và := là:
// - Biến khai báo bằng var có thể sử dụng trong hoặc ngoài hàm, trong khi biến khai báo bằng := chỉ có thể sử dụng trong hàm.
// - Khai báo và gán giá trị bằng var có thể tách rời, trong khi khai báo và gán giá trị bằng := phải thực hiện cùng lúc.
