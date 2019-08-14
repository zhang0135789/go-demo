package main

import "fmt"

/*
	形式参数
	形式参数会作为函数的局部变量来使用。
*/

/* 声明全局变量 */
var a int = 20

func main() {
	/* main 函数中声明局部变量 */
	var a int = 10
	var b int = 20
	var c int = 0

	c = sum(a, b)
	fmt.Printf("main()函数中 c = %d\n", c)

	fmt.Printf("main()函数中 a = %d\n", a)
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	a = a + 1
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}
