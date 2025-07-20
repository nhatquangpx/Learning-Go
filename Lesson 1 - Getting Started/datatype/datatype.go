// Có 4 loại kiểu dữ liệu cơ bản trong Go: string, int, bool, float64
// - string: Kiểu dữ liệu chuỗi, dùng để lưu trữ văn bản.
// - int: Kiểu dữ liệu số nguyên, dùng để lưu trữ số nguyên.
// - bool: Kiểu dữ liệu boolean, dùng để lưu trữ giá trị đúng (true) hoặc sai (false).
// - float64: Kiểu dữ liệu số thực, dùng để lưu trữ số có phần thập phân.

package main

import "fmt"

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}

// Cách lấy kiểu dữ liệu của biến trong Go:
// - Sử dụng hàm fmt.Printf với định dạng %T để in ra kiểu dữ liệu của biến.
// - Ví dụ: fmt.Printf("Type of a: %T\n", a) sẽ in ra kiểu dữ liệu của biến a.

// Cách lấy giá trị của biến trong Go:
// - Sử dụng hàm fmt.Println với định dạng %v để in ra giá trị của biến.
// - Ví dụ: fmt.Println("Value of a: %v", a) sẽ in ra giá trị của biến a.

// Số nguyên trong Go có hai loại: 
// - int: Số nguyên có dấu, kích thước phụ thuộc vào hệ thống (32-bit hoặc 64-bit).
// - uint: Số nguyên không dấu, cũng phụ thuộc vào hệ thống.

