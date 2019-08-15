package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	//demo2_1()

	demo2_2()

}

/*
结构体指针
*/
func demo2_2() {
	var Book1 Book
	var Book2 Book

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook2(&Book1)

	/* 打印 Book2 信息 */
	printBook2(&Book2)

}

/*
	结构体作为函数参数
*/
func demo2_1() {
	var Book1 Book
	/* 声明 Book1 为 Books 类型 */
	var Book2 Book
	/* 声明 Book2 为 Books 类型 */
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407
	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700
	/* 打印 Book1 信息 */
	printBook(Book1)
	/* 打印 Book2 信息 */
	printBook(Book2)
}

func printBook(book Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook2(book *Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
