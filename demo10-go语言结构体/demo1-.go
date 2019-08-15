package main

import "fmt"

/*
	Go 语言结构体
	定义结构体
	type struct_variable_type struct {
	   member definition;
	   member definition;
	   ...
	   member definition;
	}
	声明
	variable_name := structure_variable_type {value1, value2...valuen}
	或
	variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	demo1()

	demo2()

}

/*
	访问结构体成员
*/
func demo2() {
	/* 声明 Book1 Book2 为 Books 类型 */
	var Book1 Books
	var Book2 Books

	/* book 1 描述 */
	Book1.title = "go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "go jiaoasd"
	Book1.book_id = 110

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 120

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

/*
定义结构体
*/
func demo1() {
	// 创建一个新的结构体
	fmt.Println(Books{"go语言", "www.runoob.com", "hello world", 1})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "go语言", subject: "hello world", author: "www.runoob.com", book_id: 2})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
